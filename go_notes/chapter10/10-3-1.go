package main

import (
	"fmt"
	"reflect"
)

type X struct{}

func (X) Test(x, y int) (int, error) {

	return x + y, fmt.Errorf("err: %d", x+y)
}

func main() {
	var a X
	v := reflect.ValueOf(&a)
	m := v.MethodByName("Test")

	in := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	}

	out := m.Call(in)
	for _, v := range out {
		fmt.Println(v)
	}
}
