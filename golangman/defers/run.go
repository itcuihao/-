package defers

import "fmt"

// return 不是原子操作，
// a中 return i=1, defer i`++(这个i`是原i的拷贝)
// 最后 return 1
func a() (i int) {
	defer func(i int) {
		fmt.Println("a:", i)
		i++
	}(i)
	fmt.Println(i)
	return 1
}

func b() (i int) {
	defer func() {
		fmt.Println("b:", i)
		i++
	}()
	fmt.Println(i)
	return 1
}

// 对比c,d，可得知，defer 中的值是函数顺序执行时赋值
func c() (i int) {
	t := 2
	defer func() {
		fmt.Println("it:", i, t)
		i += t
	}()
	return t
}

func d() (i int) {
	t := 2
	defer func() {
		fmt.Println("it:", i, t)
		i += t
	}()
	return
}

func e() (i int) {
	defer func() {
		i += 2
	}()
	return 1
}
