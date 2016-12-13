package main

import (
	"fmt"
	"reflect"
)

type S struct{}

type T struct {
	S
}

func (S) sVal()  {}
func (*S) sPtr() {}
func (T) tVal()  {}
func (*T) tPtr() {}
func methodSet(a interface{}) {

	t := reflect.TypeOf(a)
	fmt.Println(t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m)
		fmt.Println(m.Name, m.Type)
	}
}
func main() {

	var t T
	z := reflect.TypeOf(t)
	fmt.Println(z.NumMethod())
	methodSet(&T{})
	fmt.Println("...........")
	methodSet(T{})
}
