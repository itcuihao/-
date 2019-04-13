package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

// (win) set GOMAXPROCS=1 && set GODEBUG=schedtrace=1000 && example
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	runtime.GOMAXPROCS(2)
	go func() {
		fmt.Println("hao")
	}()
	go func() {
		fmt.Println("c7")
	}()
	time.Sleep(time.Second * 1)
}
