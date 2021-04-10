package gist

import "testing"

func TestHum2Underscore(t *testing.T) {
	h := hump2Underscore("baseTile")
	t.Log(h)
}

func TestBinarySearch(t *testing.T) {
	a := []int{31}
	index := binarySearch(31, a)
	t.Log(index)
}

func TestPlist(t *testing.T) {
	plist()
}
