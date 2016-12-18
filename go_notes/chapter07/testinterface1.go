package main

type stringer interface {
	string() string
}
type tester interface {
	test()
	stringer
}

type data struct{}

func (*data) test()         { println("2") }
func (data) string() string { return "1" }
func main() {
	var d data
	var t tester = &d
	t.test()
	println(t.string())
}
