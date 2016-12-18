package main

type Ner interface {
	a()
	b(int)
	c(string) string
}

type N int

func (N) a()               {}
func (*N) b(int)           {}
func (*N) c(string) string { return "1" }
func main() {
	var n N
	var t Ner = &n
	t.a()
}
