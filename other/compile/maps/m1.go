package main

// makehmap_small实现make的创建地图（map [k] v）和
//当提示知道最多为bucketCnt时，make（map [k] v，提示）
//在编译时，需要在堆上分配映射。

// makemap实现为make创建地图创建（map [k] v，提示）。
//如果编译器已确定该映射或第一个存储桶
//可以在堆栈上创建，h和/或bucket可以是非零的。
//如果h！= nil，可以直接在h中创建地图。
//如果h.buckets！= nil，指向的桶可以用作第一个桶。

// 从汇编代码看，map初始化时 map[type]type{}，make(map[type]type),make(map[type]type,len)（len<=8），
// 直接将数据交由 map 保存，远比用指针高效。这不但减少了堆内存分配，关键还在于垃圾回收器不会扫描非指针类型 key/value 对象。
// make(map[type]type,len)（len>8），会调用 makemap 方法。
// makemap_small 需要在堆上分配，
// 如果编译器已确定该映射或第一个存储桶 makemap 会在栈上创建

// 具体使用栈与堆还要看逃逸分析的结构
// 栈性能比堆要好，makemap_small 为啥用堆呢？

// https://github.com/EDDYCJY/blog/blob/master/golang/pkg/2019-03-04-%E6%B7%B1%E5%85%A5%E7%90%86%E8%A7%A3Go-map-%E5%88%9D%E5%A7%8B%E5%8C%96%E5%92%8C%E8%AE%BF%E9%97%AE%E5%85%83%E7%B4%A0.md

// map 数据结构

// A header for a Go map.

// type hmap struct {
// 	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
// 	// Make sure this stays in sync with the compiler's definition.
// 	count     int // # live cells == size of map.  Must be first (used by len() builtin)
// 	flags     uint8
// 	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
// 	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
// 	hash0     uint32 // hash seed

// 	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
// 	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
// 	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

// 	extra *mapextra // optional fields
// }

// // mapextra holds fields that are not present on all maps.
// type mapextra struct {
// 	// If both key and value do not contain pointers and are inline, then we mark bucket
// 	// type as containing no pointers. This avoids scanning such maps.
// 	// However, bmap.overflow is a pointer. In order to keep overflow buckets
// 	// alive, we store pointers to all overflow buckets in hmap.extra.overflow and hmap.extra.oldoverflow.
// 	// overflow and oldoverflow are only used if key and value do not contain pointers.
// 	// overflow contains overflow buckets for hmap.buckets.
// 	// oldoverflow contains overflow buckets for hmap.oldbuckets.
// 	// The indirection allows to store a pointer to the slice in hiter.
// 	overflow    *[]*bmap
// 	oldoverflow *[]*bmap

// 	// nextOverflow holds a pointer to a free overflow bucket.
// 	nextOverflow *bmap
// }

// // A bucket for a Go map.
// type bmap struct {
// 	// tophash generally contains the top byte of the hash value
// 	// for each key in this bucket. If tophash[0] < minTopHash,
// 	// tophash[0] is a bucket evacuation state instead.
// 	tophash [bucketCnt]uint8
// 	// Followed by bucketCnt keys and then bucketCnt values.
// 	// NOTE: packing all the keys together and then all the values together makes the
// 	// code a bit more complicated than alternating key/value/key/value/... but it allows
// 	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
// 	// Followed by an overflow pointer.
// }

// 为什么以8区分调用 makemap_small 与 makemap?
// func makemap(t *maptype, hint int, h *hmap) *hmap {}
// 当 hint 大于 8 时，就会使用 *mapextra 做溢出桶，使用栈存储。
// 若 hint 小于 8，则存储在 buckets 桶中,实际上一个bucket就是一个bmap（	tophash [bucketCnt]uint8）
// bucketCntBits = 3; bucketCnt = 1 << bucketCntBits
// bucketCnt 等于 8
// 所以当 hint<8 时，hmap 中的 buckets 就足够存储，分配在堆上。

// tophash
// tophash 是个长度为 8 的数组，代指桶最大可容纳的键值对为 8。
// 存储每个元素 hash 值的高 8 位，如果 tophash [0] <minTopHash，则 tophash [0] 表示为迁移进度

// 在这里可以注意到，（当 hint 大于等于 8 ）第一次初始化 map 时，就会通过调用 makeBucketArray 对 buckets 进行分配。因此我们常常会说，在初始化时指定一个适当大小的容量。能够提升性能。

// map 的收缩
// map 不会收缩 “不再使用” 的空间。就算把所有键值删除，它依然保留内存空间以待后用。
// 如果一个非常大的map里面的元素很少的话，可以考虑新建一个map将老的map元素手动复制到新的map中。

// mapaccess1：返回 h[key] 的指针地址，如果键不在 map 中，将返回对应类型的零值
// mapaccess2：返回 h[key] 的指针地址，如果键不在 map 中，将返回零值和布尔值用于判断
// 执行步骤
// 判断 map 是否为 nil，长度是否为 0。若是则返回零值
// 判断当前是否并发读写 map，若是则抛出异常
// 根据 key 的不同类型调用不同的 hash 方法计算得出 hash 值
// 确定 key 在哪一个 bucket 中，并得到其位置
// 判断是否正在发生扩容（h.oldbuckets 是否为 nil），若正在扩容，则到老的 buckets 中查找（因为 buckets 中可能还没有值，搬迁未完成），若该 bucket 已经搬迁完毕。则到 buckets 中继续查找
// 计算 hash 的 tophash 值（高八位）
// 根据计算出来的 tophash，依次循环对比 buckets 的 tophash 值（快速试错）
// 如果 tophash 匹配成功，则计算 key 的所在位置，正式完整的对比两个 key 是否一致
// 若查找成功并返回，若不存在，则返回零值

const (
	// Maximum number of key/value pairs a bucket can hold.
	//桶可容纳的最大键/值对数。
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits
)

// 关于预设容量
// 这个比较容易理解，map是基于hash算法实现的，通过计算key的hash值来分布和查找对象，
// 如果key的hash值相同的话，一般会通过拉链法解决冲突（Java）。如果容量太小，冲突就比较严重。
// 数据查询速度难免降低；如果需要提供数据查询速度，需要以空间换时间，加大容量。
// 如果初始容量太小，而你需要存入大量的数据，一定就会发生数据复制和rehash（很有可能发生多次，go map 的负载因子是：6.5
// Maximum average load of a bucket that triggers growth.loadFactor =6.5）。
// 所以预估容量就比较重要了。既能减少空间浪费，同时能避免运行时多次内存复制和rehash。

func main() {
	m := make(map[int]int, 9)
	m[1] = 1
}

// 0x0028 00040 (.\m1.go:6)        PCDATA  $2, $1
// 0x0028 00040 (.\m1.go:6)        PCDATA  $0, $0
// 0x0028 00040 (.\m1.go:6)        LEAQ    type.map[int]int(SB), AX
// 0x002f 00047 (.\m1.go:6)        PCDATA  $2, $0
// 0x002f 00047 (.\m1.go:6)        MOVQ    AX, (SP)
// 0x0033 00051 (.\m1.go:6)        MOVQ    $9, 8(SP)
// 0x003c 00060 (.\m1.go:6)        MOVQ    $0, 16(SP)
// 0x0045 00069 (.\m1.go:6)        CALL    runtime.makemap(SB)
// 0x004a 00074 (.\m1.go:6)        PCDATA  $2, $1
// 0x004a 00074 (.\m1.go:6)        MOVQ    24(SP), AX
// 0x004f 00079 (.\m1.go:6)        PCDATA  $0, $1
// 0x004f 00079 (.\m1.go:6)        MOVQ    AX, ""..autotmp_18+64(SP)
// 0x0054 00084 (.\m1.go:7)        PCDATA  $2, $2
// 0x0054 00084 (.\m1.go:7)        LEAQ    type.map[int]int(SB), CX
// 0x005b 00091 (.\m1.go:7)        PCDATA  $2, $1
// 0x005b 00091 (.\m1.go:7)        MOVQ    CX, (SP)
// 0x005f 00095 (.\m1.go:7)        PCDATA  $2, $0
// 0x005f 00095 (.\m1.go:7)        MOVQ    AX, 8(SP)
// 0x0064 00100 (.\m1.go:7)        MOVQ    $1, 16(SP)
// 0x006d 00109 (.\m1.go:7)        CALL    runtime.mapassign_fast64(SB)
// 0x0072 00114 (.\m1.go:7)        PCDATA  $2, $1
// 0x0072 00114 (.\m1.go:7)        MOVQ    24(SP), AX
// 0x0077 00119 (.\m1.go:7)        PCDATA  $2, $0
// 0x0077 00119 (.\m1.go:7)        MOVQ    $1, (AX)

// go tool compile -m -S .\m1.go
// func main() {
// 	m := make(map[int]int, 3)
// 	m[1] = 1
// 	// 会导致m 逃逸到堆
// 	// fmt.Println(m)
// 	println(m)
// }

// func main() {
// 	m := make(map[int]int, 8)
// 	m[1] = 1
// 	fmt.Println(m)
// }

// 0x0028 00040 (.\m1.go:61)       PCDATA  $2, $0
// 0x0028 00040 (.\m1.go:61)       PCDATA  $0, $0
// 0x0028 00040 (.\m1.go:61)       CALL    runtime.makemap_small(SB)
// 0x002d 00045 (.\m1.go:61)       PCDATA  $2, $1
// 0x002d 00045 (.\m1.go:61)       MOVQ    (SP), AX
// 0x0031 00049 (.\m1.go:61)       PCDATA  $0, $1
// 0x0031 00049 (.\m1.go:61)       MOVQ    AX, ""..autotmp_18+64(SP)
// 0x0036 00054 (.\m1.go:62)       PCDATA  $2, $2
// 0x0036 00054 (.\m1.go:62)       LEAQ    type.map[int]int(SB), CX
// 0x003d 00061 (.\m1.go:62)       PCDATA  $2, $1
// 0x003d 00061 (.\m1.go:62)       MOVQ    CX, (SP)
// 0x0041 00065 (.\m1.go:62)       PCDATA  $2, $0
// 0x0041 00065 (.\m1.go:62)       MOVQ    AX, 8(SP)
// 0x0046 00070 (.\m1.go:62)       MOVQ    $1, 16(SP)
// 0x004f 00079 (.\m1.go:62)       CALL    runtime.mapassign_fast64(SB)
// 0x0054 00084 (.\m1.go:62)       PCDATA  $2, $1
// 0x0054 00084 (.\m1.go:62)       MOVQ    24(SP), AX
// 0x0059 00089 (.\m1.go:62)       PCDATA  $2, $0
// 0x0059 00089 (.\m1.go:62)       MOVQ    $1, (AX)

// func main() {
// 	m := map[int]int{}
// 	m[1] = 1
// 	fmt.Println(m)
// }

// 0x0028 00040 (.\m1.go:6)        PCDATA  $2, $0
// 0x0028 00040 (.\m1.go:6)        PCDATA  $0, $0
// 0x0028 00040 (.\m1.go:6)        CALL    runtime.makemap_small(SB)
// 0x002d 00045 (.\m1.go:6)        PCDATA  $2, $1
// 0x002d 00045 (.\m1.go:6)        MOVQ    (SP), AX
// 0x0031 00049 (.\m1.go:6)        PCDATA  $0, $1
// 0x0031 00049 (.\m1.go:6)        MOVQ    AX, ""..autotmp_18+64(SP)
// 0x0036 00054 (.\m1.go:7)        PCDATA  $2, $2
// 0x0036 00054 (.\m1.go:7)        LEAQ    type.map[int]int(SB), CX
// 0x003d 00061 (.\m1.go:7)        PCDATA  $2, $1
// 0x003d 00061 (.\m1.go:7)        MOVQ    CX, (SP)
// 0x0041 00065 (.\m1.go:7)        PCDATA  $2, $0
// 0x0041 00065 (.\m1.go:7)        MOVQ    AX, 8(SP)
// 0x0046 00070 (.\m1.go:7)        MOVQ    $1, 16(SP)
// 0x004f 00079 (.\m1.go:7)        CALL    runtime.mapassign_fast64(SB)
// 0x0054 00084 (.\m1.go:7)        PCDATA  $2, $1
// 0x0054 00084 (.\m1.go:7)        MOVQ    24(SP), AX
// 0x0059 00089 (.\m1.go:7)        PCDATA  $2, $0
// 0x0059 00089 (.\m1.go:7)        MOVQ    $1, (AX)

// func main() {
// 	m := make(map[int]int)
// 	m[1] = 1
// 	fmt.Println(m)
// }

// 0x0028 00040 (.\m1.go:6)        PCDATA  $2, $0
// 0x0028 00040 (.\m1.go:6)        PCDATA  $0, $0
// 0x0028 00040 (.\m1.go:6)        CALL    runtime.makemap_small(SB)
// 0x002d 00045 (.\m1.go:6)        PCDATA  $2, $1
// 0x002d 00045 (.\m1.go:6)        MOVQ    (SP), AX
// 0x0031 00049 (.\m1.go:6)        PCDATA  $0, $1
// 0x0031 00049 (.\m1.go:6)        MOVQ    AX, ""..autotmp_18+64(SP)
// 0x0036 00054 (.\m1.go:7)        PCDATA  $2, $2
// 0x0036 00054 (.\m1.go:7)        LEAQ    type.map[int]int(SB), CX
// 0x003d 00061 (.\m1.go:7)        PCDATA  $2, $1
// 0x003d 00061 (.\m1.go:7)        MOVQ    CX, (SP)
// 0x0041 00065 (.\m1.go:7)        PCDATA  $2, $0
// 0x0041 00065 (.\m1.go:7)        MOVQ    AX, 8(SP)
// 0x0046 00070 (.\m1.go:7)        MOVQ    $1, 16(SP)
// 0x004f 00079 (.\m1.go:7)        CALL    runtime.mapassign_fast64(SB)
// 0x0054 00084 (.\m1.go:7)        PCDATA  $2, $1
// 0x0054 00084 (.\m1.go:7)        MOVQ    24(SP), AX
// 0x0059 00089 (.\m1.go:7)        PCDATA  $2, $0
// 0x0059 00089 (.\m1.go:7)        MOVQ    $1, (AX)

// func main() {
// 	m := make(map[int]int, 5)
// 	m[1] = 1
// 	fmt.Println(m)
// }
// 0x0028 00040 (.\m1.go:6)        PCDATA  $2, $0
// 0x0028 00040 (.\m1.go:6)        PCDATA  $0, $0
// 0x0028 00040 (.\m1.go:6)        CALL    runtime.makemap_small(SB)
// 0x002d 00045 (.\m1.go:6)        PCDATA  $2, $1
// 0x002d 00045 (.\m1.go:6)        MOVQ    (SP), AX
// 0x0031 00049 (.\m1.go:6)        PCDATA  $0, $1
// 0x0031 00049 (.\m1.go:6)        MOVQ    AX, ""..autotmp_18+64(SP)
// 0x0036 00054 (.\m1.go:7)        PCDATA  $2, $2
// 0x0036 00054 (.\m1.go:7)        LEAQ    type.map[int]int(SB), CX
// 0x003d 00061 (.\m1.go:7)        PCDATA  $2, $1
// 0x003d 00061 (.\m1.go:7)        MOVQ    CX, (SP)
// 0x0041 00065 (.\m1.go:7)        PCDATA  $2, $0
// 0x0041 00065 (.\m1.go:7)        MOVQ    AX, 8(SP)
// 0x0046 00070 (.\m1.go:7)        MOVQ    $1, 16(SP)
// 0x004f 00079 (.\m1.go:7)        CALL    runtime.mapassign_fast64(SB)
// 0x0054 00084 (.\m1.go:7)        PCDATA  $2, $1
// 0x0054 00084 (.\m1.go:7)        MOVQ    24(SP), AX
// 0x0059 00089 (.\m1.go:7)        PCDATA  $2, $0
// 0x0059 00089 (.\m1.go:7)        MOVQ    $1, (AX)
