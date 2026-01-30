package iteration

import "strings"

func Repeat(char string, n int) string {
	var repeated strings.Builder
	for range n {
		repeated.WriteString(char)
	}

	return repeated.String()
}
