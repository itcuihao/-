package main

import "fmt"

type Animal interface {
	Grow()
	Move(string) string
}

type Cat struct {
	Name    string
	Age     int32
	Address string
}

func (cat *Cat) Grow() {
	cat.Age++
}

func (cat *Cat) Move(newAddress string) (oldAddress string) {
	oldAddress = cat.Address
	cat.Address = newAddress
	return
}

func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	animal, ok := interface{}(&myCat).(Animal)
	animal.Grow()
	fmt.Printf("%v, %v\n", ok, animal)
}
