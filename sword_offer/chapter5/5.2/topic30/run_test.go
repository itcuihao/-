package topic30

import "testing"

func TestRun(t *testing.T) {
	// numbers := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	numbers := []int{3, 2, 1, 1}
	i := Run1(numbers, 2)
	t.Log(numbers)
	t.Log(i)
}
