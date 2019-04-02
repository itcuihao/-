package topic30

// 时间复杂度为O(kn)。缺点就是，时间复杂度大了点，关键是改变了原有的输入数组。
func Run1(numbers []int, k int) []int {
	list := make([]int, 0, k)
	if numbers == nil || len(numbers) < k {
		return list
	}

	for i := 0; i < k; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	list = numbers[:k]
	return list
}
