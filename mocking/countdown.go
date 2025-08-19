package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countDownStart = 3

func Countdown(writer io.Writer) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
