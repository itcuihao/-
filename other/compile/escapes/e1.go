package main

import "fmt"

func main() {

	a := "a"
	b := i(a)
	c := i(a)
	fmt.Println(b == c)
	fmt.Printf("%x %P\n", b, b)
	fmt.Printf("%x %P\n", c, c)

	u := User{
		Name: "a",
	}
}

func i(s string) interface{} {
	var in interface{}
	in = s
	return in
}

type User struct {
	A
	Age int
}
type A struct {
	Name string
}

func F() A {
	a := new(A)
	a.Name = "1"
	a = nil

	return *a
}
