package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for tl := 1; tl <= taskLoad; tl++ {
		tasks <- fmt.Sprintf("Task : %d", tl)
	}

	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		sleep := rand.Int63n(100)
		fmt.Println("s::", sleep)
		fmt.Println("sd::", time.Duration(sleep))
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
