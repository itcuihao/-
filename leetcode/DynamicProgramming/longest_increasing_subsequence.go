package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(LogestIncreasingSubsequence(nums))
}

func LogestIncreasingSubsequence(nums []int) (nlen int) {
	tmp := make([]int, len(nums))
	for k, v := range nums {
		tmp[k] = 1
		for i := 0; i < k; i++ {
			fmt.Println("i:", nums[i])
			fmt.Println("k:", nums[k])

			if nums[i] < nums[k] {
				fmt.Println("ti:", tmp[i])
				fmt.Println("tk:", tmp[k])
				if tmp[k] > tmp[i]+1 {
					tmp[k] = v
				}
				tmp[k] = tmp[i] + 1
			}
		}
		if tmp[k] > nlen {
			nlen = tmp[k]
			fmt.Println("len:", nlen)
		}
		fmt.Println("---")
	}
	fmt.Println(tmp)
	return
}
