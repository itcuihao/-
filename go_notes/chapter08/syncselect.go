package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	a, b := make(chan int), make(chan int)
	go func() {
		defer wg.Done()
		for {
			var (
				name string
				x    int
				ok   bool
			)
			select {
			case x, ok = <-a:
				println("name a", x)
				if !ok {
					println("!ok-a")
					a = nil
					break
				}
				name = "a"
			case x, ok = <-b:
				println("name b", x)
				if !ok {
					println("!ok-b")
					b = nil //将通道设置为nil，这样就会阻塞，不会被select选中
					break
				}
				name = "b"
			}
			if a == nil || b == nil {
				println("!ok")
				return
			}
			println(name, x)
		}
	}()
	go func() {
		defer wg.Done()
		defer close(a)
		for i := 0; i < 5; i++ {
			select {
			case a <- i:
				println("aaa", i)

			}
		}
	}()
	go func() {
		defer wg.Done()
		defer close(b)
		for i := 0; i < 5; i++ {
			select {
			case b <- i * 10:
				println("bbb", i)
			}
		}
	}()
	wg.Wait()
}
