package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	lisa.changeEmail("lisa@qq.com")
	lisa.notify()

	bill.changeEmail("bill@qq.com")
	bill.notify()

}
