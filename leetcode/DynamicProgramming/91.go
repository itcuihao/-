package main

import (
	"fmt"
)

func main() {
	fmt.Println("ways:", decodeways("26"))
}

func decodeways(s string) (n int) {
	println(s)
	println(s[0:1])
	sl := len(s)
	if sl == 0 {
		return
	}
	m := make([]int, sl+1)
	m[sl] = 1
	if s[sl-1] == '0' {
		m[sl-1] = 0
	} else {
		m[sl-1] = 1
	}
	for i := sl - 2; i >= 0; i-- {
		println("i:", i)
		fmt.Println("m:", m)
		if s[i] == '0' {
			continue
		}
		println("s[i:i+2]:", s[i:i+2])
		if s[i:i+2] <= "26" {
			println("m[i+1]:", i+1, m[i+1])
			println("m[i+2]:", i+2, m[i+2])
			println("mi:", i, m[i])
			m[i] = m[i+1] + m[i+2]
			println("mi:", i, m[i])
		} else {
			m[i] = m[i+1]
			println("mi:", i, m[i])
		}
	}
	n = m[0]
	return
}
