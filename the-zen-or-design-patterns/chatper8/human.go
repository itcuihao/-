package chatper8

import "fmt"

type Human interface {
	getColor() string
	talk()
}

type Black struct {
}

func (h Black) getColor() string {
	return "black"
}

func (h Black) talk() {
	fmt.Println("black")
}

type Yellow struct {
}

func (h Yellow) getColor() string {
	return "yellow"
}

func (h Yellow) talk() {
	fmt.Println("yellow")
}

type White struct {
}

func (h White) getColor() string {
	return "white"
}

func (h White) talk() {
	fmt.Println("white")
}

func createHuman(h Human) {
	c := h.getColor()
	fmt.Println("color: ", c)
	h.talk()
}

func NvWa() {
	h := &Black{}
	createHuman(h)
}
