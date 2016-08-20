package array

func RemoveDuplicates(A []int, n int) (j int) {
	if n == 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		if A[j] != A[i] {
			j++
			A[j] = A[i]
		}
	}
	return j + 1
}
