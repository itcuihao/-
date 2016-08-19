package main

import (
	"fmt"
)

func main() {

	data := map[string]string{"haha": "123"}

	fmt.Println(data)
	fmt.Print(data["haha"])
}
