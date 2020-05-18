package chapter09

import (
	"testing"
	"time"
)

func TestTeller(t *testing.T) {
	InitTeller()
	go func() {
		Deposit(100)
		balance := Balance()
		t.Log(balance)
	}()
	go func() {
		Deposit(200)
		balance := Balance()
		t.Log(balance)
	}()
	time.Sleep(3 * time.Second)
}
