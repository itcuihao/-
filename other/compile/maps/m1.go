package main

import "fmt"

// makehmap_small实现make的创建地图（map [k] v）和
//当提示知道最多为bucketCnt时，make（map [k] v，提示）
//在编译时，需要在堆上分配映射。

// makemap实现为make创建地图创建（map [k] v，提示）。
//如果编译器已确定该映射或第一个存储桶
//可以在堆栈上创建，h和/或bucket可以是非零的。
//如果h！= nil，可以直接在h中创建地图。
//如果h.buckets！= nil，指向的桶可以用作第一个桶。

// 从汇编代码看，map初始化时 map[type]type{}，make(map[type]type),make(map[type]type,len)（len<=8），会调用 makemap_small 方法，
// make(map[type]type,len)（len>8），会调用 makemap 方法。
// makemap_small 需要在堆上分配，
// 如果编译器已确定该映射或第一个存储桶 makemap 会在栈上创建

// 栈性能比堆要好，makemap_small 为啥不用堆呢？

const (
	// Maximum number of key/value pairs a bucket can hold.
	//桶可容纳的最大键/值对数。
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits
)

// func main() {
// 	m := make(map[int]int, 9)
// 	m[1] = 1
// 	fmt.Println(m)

// 	fmt.Println(bucketCnt)
// }

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

func main() {
	m := make(map[int]int, 8)
	m[1] = 1
	fmt.Println(m)
}

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
