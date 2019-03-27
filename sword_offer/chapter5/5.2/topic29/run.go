package topic29

import "fmt"

// O(n)
func MoreThanHalfNum(numbers []int) int {
	length := len(numbers)
	if CheckInvalidArray(numbers, length) {
		return 0
	}

	middle := length >> 1
	start, end := 0, length-1

	index := Partition(numbers, start, end)

	for index != middle {
		if index > middle {
			end = index - 1
			index = Partition(numbers, start, end)
		} else {
			start = index + 1
			index = Partition(numbers, start, end)
		}
	}
	result := numbers[middle]
	if !CheckMoreThanHalf(numbers, result) {
		result = 0
	}
	return result
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

	return start
}

func CheckInvalidArray(numbers []int, length int) bool {
	return numbers == nil || length <= 0
}

func CheckMoreThanHalf(numbers []int, number int) bool {
	t := 0
	for _, v := range numbers {
		if v == number {
			t++
		}
	}
	return t*2 > len(numbers)
}
func Quick(num []int, start, end int) {
	if start < end {
		base := Partition(num, start, end)
		fmt.Println(base, start, end)
		Quick(num, start, base-1)
		Quick(num, base+1, end)
	}
}

// O(n)
func MoreThanHalfNum2(numbers []int) int {
	l := len(numbers)
	if CheckInvalidArray(numbers, l) {
		return 0
	}

	result := numbers[0]
	times := 1

	for i := 0; i < l; i++ {
		if times == 0 {
			result = numbers[i]
			times = 1
		} else if numbers[i] == result {
			times++
		} else {
			times--
		}
	}
	if !CheckMoreThanHalf(numbers, result) {
		result = 0
	}
	return result
}
