package main

import (
	"fmt"
	"strings"
)

func rm_DuplicateSlice(list *[]string) []string {
	x := []string{}
	for _, i := range *list {
		if len(x) == 0 {

			x = append(x, i)
			fmt.Println("1====", i)
			fmt.Println("1====", x)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
					fmt.Println("2====", i)
					fmt.Println("2====", x)
				}
			}
		}
	}
	return x
}
func main() {
	var slicestr []string
	fmt.Println(1)
	str1 := "11011101"
	fmt.Println(str1[:1])
	slicestr = append(slicestr, str1[:1])
	fmt.Println(str1[:2])
	slicestr = append(slicestr, str1[:2])
	fmt.Println(str1[:4])
	slicestr = append(slicestr, str1[:4])
	fmt.Println(str1[:5])
	slicestr = append(slicestr, str1[:5])
	slicestr = append(slicestr, str1)

	str2 := "11011102"
	fmt.Println(str2[:1])
	slicestr = append(slicestr, str2[:1])
	fmt.Println(str2[:2])
	slicestr = append(slicestr, str2[:2])
	fmt.Println(str2[:4])
	slicestr = append(slicestr, str2[:4])
	fmt.Println(str2[:5])
	slicestr = append(slicestr, str2[:5])
	slicestr = append(slicestr, str2)

	fmt.Println(slicestr)
	single_slicestr := rm_DuplicateSlice(&slicestr)
	fmt.Println(single_slicestr)
	fmt.Println(strings.Join(single_slicestr, ","))
}
