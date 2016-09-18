package main

import "fmt"

var a, b, c, d = c + b, f(), f(), 3

func f() int {
	d++
	return d
}

func main() {

	fmt.Println(a, b, c, d)

	w, x, y, z := 1, 1, 3, 2
	switch x {
	case w, y:
		fmt.Println("w|y")
		fallthrough
	case z:
		fmt.Println("z")
	default:
		fmt.Println("xyz")
	}
	fmt.Println("………………")
	data := [3]string{"a", "b", "c"}
	for i := range data {
		fmt.Println(i, data[i])
	}
	for _, s := range data {
		fmt.Println(s)
	}
}
