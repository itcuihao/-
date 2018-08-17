package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
	ss      []int
)

func main() {
	wg.Add(4)
	go incCounter(1)
	go incCounter(2)
	go setSlice()
	go setSlice()
	wg.Wait()

	fmt.Println("Final counter:", counter)
	fmt.Println("Final ss:", ss)
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		mutex.Lock()
		value := counter
		runtime.Gosched()
		value++
		counter = value
		mutex.Unlock()
	}
}

func setSlice() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		ss = append(ss, i)
		mutex.Unlock()
	}
}
