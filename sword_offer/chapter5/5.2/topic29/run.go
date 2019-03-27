package topic29

import "fmt"

func MonrThanHalfNum(numbers []int, length int) int {
	if CheckInvalidArray(numbers, length) {
		return 0
	}

	// middle := length >> 1
	// start, end := 0, length-1

	return 0
}

func Partition(data []int, start, end int) int {
	if data == nil || start < 0 || end >= len(data) {
		return -1
	}

	v := data[start]

	for start < end {
		for start < end && data[start] < v {
			start++
		}
		fmt.Println(start)
		for start < end && data[end] > v {
			end--
		}
		data[start], data[end] = data[end], data[start]
	}

	data[start] = v

	fmt.Println(data)
	return start
}

func CheckInvalidArray(numbers []int, length int) bool {
	return numbers == nil && length <= 0
}

func Quick(num []int, start, end int) {
	if start < end {
		base := Partition(num, start, end)
		fmt.Println(base, start, end)
		Quick(num, start, base-1)
		Quick(num, base+1, end)
	}
}
