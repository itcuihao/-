package gist

import (
	"strings"
	"unicode"
)

// 驼峰转下划线
func hump2Underscore(name string) string {
	builder := strings.Builder{}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				builder.WriteByte('_')
			}
			builder.WriteRune(unicode.ToLower(r))
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func binarySearch(target int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != target {
		return false
	}

	return true
}
