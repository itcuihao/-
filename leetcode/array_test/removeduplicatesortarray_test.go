package array

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	s := []int{1, 2, 2, 3, 4, 4, 5}
	l := RemoveDuplicates(s, 7)
	fmt.Println(l)
	if l != len(s)-2 {
		t.Error(l)
		return
	}
}
