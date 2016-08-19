package main

import (
	"fmt"
)

//消费者
func consumer(data chan int, done chan bool) {
	for x := range data {
		println("recv:", x)
		fmt.Println(data)

	}
	fmt.Println("333")
	fmt.Println(data)

	done <- true
	fmt.Println("444")
	fmt.Println(data)

}

//生产者
func producer(data chan int) {
	for i := 0; i < 4; i++ {

		data <- i
	}

	fmt.Println("111")
	fmt.Println(data)

	close(data)
	fmt.Println("22")
	fmt.Println(data)

}

func main() {
	done := make(chan bool)
	data := make(chan int)
	go producer(data)
	go consumer(data, done)

	<-done
}
