package iteration

func Repeat(char string, n int) string {
	var repeated string
	for range n {
		repeated += char
	}

	return repeated
}
