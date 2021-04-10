package chatper13

import "fmt"

func testBubbleSort() {
	nums := []int{5, 3, 4, 1, 2}
	r := bubbleSort(nums)
	fmt.Println(r)
}

// 冒泡排序
func bubbleSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		fmt.Println(nums)
	}
	return nums
}

func testInsertSort() {
	nums := []int{5, 3, 4, 1, 2}
	r := insertSort(nums)
	fmt.Println(r)
}

// 插入排序
func insertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > temp {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = temp
	}
	return nums
}

func testMergeSort() {
	nums := []int{5, 3, 4, 1, 2}
	r := customMergeSort(nums)
	fmt.Println(r)
}

func customMergeSort(nums []int) []int {
	s, e := 0, len(nums)
	if s == e {
		return nums
	}
	mergeSort(nums, s, e)
	return nums
}

func mergeSort(nums []int, s, e int) {
	if s == e {
		return
	}
	m := (s + e) / 2
	mergeSort(nums, s, m)
	mergeSort(nums, m+1, e)
	merge(nums, s, m, e)
}

func merge(nums []int, s, m, e int) {
	left := m - s + 1
	right := e - m
	leftNums := make([]int, left)
	rightNums := make([]int, right)
	for i := 0; i < left; i++ {
		leftNums[i] = nums[s+i]
	}
	for i := 0; i < right; i++ {
		rightNums[i] = nums[m+i+1]
	}
	i, j, k := 0, 0, s
	for ; k <= e && i < left && j < right; k++ {
		if leftNums[i] <= rightNums[j] {
			nums[k] = leftNums[i]
			i++
		} else {
			nums[k] = rightNums[j]
			j++
		}
	}
	for ; k <= e && i < left; k++ {
		nums[k] = leftNums[i]
		i++
	}
	for ; k <= e && j < right; k++ {
		nums[k] = rightNums[j]
		j++
	}
}
