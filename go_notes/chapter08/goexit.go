package main

import (
	"runtime"
)

func main() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		defer println("a")

		func() {
			defer func() {
				println("b", recover() == nil)
			}()
			func() {
				println("c")

				runtime.Goexit()

				println("c clone")
			}()
			println("b clone")
		}()
		println("a clone")
	}()
	<-exit
	println("main exit.")
}
