package main

import (
	"fmt"
)

type N int

func (n N) test() {
	fmt.Printf("test:n, %p, %v\n", &n, n)
}
func (n *N) test2() {
	fmt.Printf("test:n, %p, %v\n", n, *n)
}
func call(m func()) {
	m()
}

func main() {
	var n N = 100
	p := &n
	n++
	f1 := n.test
	n++
	f2 := p.test
	f3 := p.test2
	f3()
	n++
	fmt.Printf("main.n: %p, %v\n", p, n)
	f1()
	f2()
	call(f1)
}
