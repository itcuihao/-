package main

import (
	"fmt"
)

func main() {
	arrayX := []int{1, 2, 2, 2, 3}
	fmt.Println("arrayX:", arrayX)
	l := RemoveElement(arrayX, 5, 2)
	fmt.Println("remove element length arrayX:", l)
}

func RemoveElement(A []int, n int, element int) int {
	j := 0
	for i := 0; i < n; i++ {
		if A[i] == element {
			continue
		}
		A[j] = A[i]
		j++
	}
	return j
}
