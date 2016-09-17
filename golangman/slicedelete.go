package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	var d []int
	fmt.Println(s)
	for i := 0; i < 6; i++ {
		d = append(s[:0], s[i:]...)
	}
	fmt.Println(d)
	fmt.Println("-----")
	fmt.Println(s)
	for k, v := range s {
		if v == 3 {
			s = append(s[:k], s[k+1:]...)
		}
	}
	fmt.Println(s)
}
