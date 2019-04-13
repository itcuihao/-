package main

import (
	"fmt"
	"runtime"
	"time"
)

// (win) set GOMAXPROCS=1 && set GODEBUG=schedtrace=1000 && example
func main() {

	runtime.GOMAXPROCS(1)
	go func() {
		for {
			fmt.Println("hao")
			time.Sleep(time.Microsecond * 100)
			break
		}
	}()
	go func() {
		fmt.Println("c7")
	}()
	time.Sleep(time.Second * 1)
}
