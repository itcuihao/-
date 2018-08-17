package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(a[len(a)/2:], c)
	go sum(a[:len(a)/2], c)

	x, y := <-c, <-class
	fmt.Println(x, "+", y, "=", x+y)
}
