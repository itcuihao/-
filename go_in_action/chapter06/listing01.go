package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	fmt.Printf("CPU NUM:%d\n", runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("wating to finish")
	wg.Wait()
	fmt.Println("\nterminating program")
}
