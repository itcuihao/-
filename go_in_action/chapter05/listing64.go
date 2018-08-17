package main

import (
	"GoInAction/chapter5/listing64/counters"
	"fmt"
)

func main() {
	// counter1 := counters.alertCounter(10)
	counter := counters.New(10)
	fmt.Println(counter)
}
