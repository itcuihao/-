package size

import (
	"math"
	"testing"
)

var a = 1
var b = 2

func TestAdd(t *testing.T) {
	c := add(a, b)
	t.Log(c)
}
func TestMul(t *testing.T) {
	a = math.MaxInt64
	b = math.MaxInt64
	c := mul(a, b)
	t.Log(c)
}
