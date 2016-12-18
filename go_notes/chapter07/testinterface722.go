package main

import (
	"log"
)

type TestError struct{}

func (*TestError) Error() string {
	return "error"
}

func test(x int) (int, error) {
	//	var err *TestError

	if x < 0 {
		println(new(*TestError))
		return 0, new(TestError)
	}
	return x + 100, nil
}

func main() {
	x, err := test(-100)
	if err != nil {
		log.Fatalln("err!=nil")
	}
	println(x)
}
