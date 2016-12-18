package main

type data struct{}

func (data) string() string {
	return "1"
}

type node struct {
	data interface {
		string() string
	}
}

func main() {
	var t interface {
		string() string
	} = data{}
	n := node{
		data: t,
	}
	println(n.data.string())
}
