package topic29

import "testing"

func TestMoreThanHalfNum(t *testing.T) {
	// numbers := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	numbers := []int{3, 2, 1, 1}
	i := MoreThanHalfNum(numbers)
	t.Log(numbers)
	t.Log(i)
}
func TestMoreThanHalfNum2(t *testing.T) {
	numbers := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	// numbers := []int{3, 2, 1, 1}
	i := MoreThanHalfNum2(numbers)
	t.Log(numbers)
	t.Log(i)
}
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
