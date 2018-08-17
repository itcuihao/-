package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Printf("sending admin email to %s<%s>\n", a.name, a.email)
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

func sendNotification(n notifier) {
	n.notify()
}
