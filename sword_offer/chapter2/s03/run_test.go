package s03

import "testing"

func TestRun(t *testing.T) {
	a := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	t.Log(run(a, 5))
}
