package main

import (
	"fmt"
)

func main() {

	data := map[string]string{"haha": "123"}

	fmt.Println(data)
	fmt.Println(data["haha"])
	str := "abcde$@"
	strbyte := []byte(str)
	for k, v := range strbyte {
		fmt.Print(k)
		fmt.Println(string(v))
	}
	fmt.Println(str[2:3])
}
