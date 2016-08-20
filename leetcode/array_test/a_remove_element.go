package array

func A_RemoveElement(A []int, n int, element int) int {
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
