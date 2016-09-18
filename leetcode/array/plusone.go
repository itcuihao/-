package main

import (
	"fmt"
)

func main() {
	fmt.Println("start...")

	s := []int{1, 2}
	fmt.Println(s)
	s2 := plusOne(s)
	fmt.Println(s2)
	fmt.Println("start...")
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}
	n := len(digits)
	ans := make([]int, n+1)
	c := 1
	for i := n - 1; i >= 0; i-- {
		ans[i+1] = (c + digits[i]) % 10
		c = (c + digits[i]) / 10
	}
	if c == 1 {
		ans[0] = 1
		return ans
	} else {
		for i := 0; i < n; i++ {
			digits[i] = ans[i+1]
		}
		return digits
	}
}
