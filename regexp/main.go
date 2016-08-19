package main

import (
	"fmt"
	"regexp"
)

func main() {
	m, err := regexp.MatchString(`[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{4}`, "370784199301042010")
	matched, err1 := regexp.MatchString(`[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{4}`, "372926")
	fmt.Println("---->")
	fmt.Println(m)
	fmt.Println(err)
	fmt.Println(matched)
	fmt.Println(err1)

}
