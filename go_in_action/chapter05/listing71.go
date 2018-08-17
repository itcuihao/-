package main

import (
	"GoInAction/chapter5/listing71/entities"
	"fmt"
)

func main() {
	// u := entities.User{
	// 	Name:  "Bill",
	// 	email: "bill@email.com",
	// }

	a := entities.Admin{
		Rights: 10,
	}
	a.Name = "Bill"
	// a.Email = "bill@email.com"
	fmt.Println(a)
	a.SetEmail("bill@email.com")

	fmt.Println(a)
}
