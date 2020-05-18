package gist

import "testing"

func TestHum2Underscore(t *testing.T) {
	h := hump2Underscore("baseTile")
	t.Log(h)
}
