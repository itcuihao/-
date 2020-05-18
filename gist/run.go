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
