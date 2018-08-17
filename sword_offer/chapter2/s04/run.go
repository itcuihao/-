package s04

import "fmt"

func run(s string) (ss string) {
	for _, v := range s {
		if v == 32 {
			ss += "%20"
			continue
		}
		ss += fmt.Sprintf("%s", string(v))
	}
	return
}
