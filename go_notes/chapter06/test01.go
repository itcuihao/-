package main

import (
	"fmt"
)

type N int

func (n N) toString() string {
	return fmt.Sprintf("%#x", n)
}
func main() {
	var a N = 22
	fmt.Println(a.toString())
}
