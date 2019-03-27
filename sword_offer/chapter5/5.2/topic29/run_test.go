package topic29

import "testing"

func TestPartition(t *testing.T) {
	// numbers := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	numbers := []int{3, 2, 1, 1}
	i := Partition(numbers, 0, len(numbers)-1)
	t.Log(numbers)
	t.Log(i)
}
func TestQuick(t *testing.T) {
	// numbers := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	numbers := []int{3, 2, 1}
	Quick(numbers, 0, len(numbers)-1)
	t.Log(numbers)
}
