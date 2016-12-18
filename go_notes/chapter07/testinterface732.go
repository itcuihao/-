package main

import (
	"fmt"
)

func main() {
	var x interface{} = func(x int) string {
		return fmt.Sprintf("d: %d", x)
	}
	switch v := x.(type) {
	case nil:
		println("nil")
	case *int:
		println(*v)
	case func(int) string:
		fmt.Println(v(100))
	default:
		println("unknow")
	}
}
