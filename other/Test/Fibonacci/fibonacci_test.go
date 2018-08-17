package Fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	r := Fibonacci(10)
	if r == 55 {
		t.Error("Fibonacci 通过")
	} else {
		t.Log("测试不通过")
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}
