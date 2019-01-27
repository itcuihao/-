package topic12

import (
	"fmt"
	"testing"
)

func TestF1(t *testing.T) {
	f1(2)
}
func TestF2(t *testing.T) {
	f2(2)
}
func TestIncrement(t *testing.T) {
	s := new(Strn)
	s.Number = []rune{'9'}
	b := s.increment()
	for _, v := range s.Number {
		fmt.Println(string(v))
	}
	fmt.Println(b)

	// printn(n)
}

func TestPrintn(t *testing.T) {
	n := []rune{'1', '0'}
	printn(n)
}
