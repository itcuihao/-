package array

import (
	"fmt"
	"testing"
)

func TestA_RemoveElement(t *testing.T) {
	s := []int{1, 2, 2, 3, 4, 2, 5, 6, 3}
	l := A_RemoveElement(s, 9, 2)
	fmt.Println(l)
	if l != len(s)-3 {
		t.Error(l)
		return
	}
}
