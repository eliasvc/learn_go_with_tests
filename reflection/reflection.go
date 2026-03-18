package main

func walk(x any, fn func(string)) {
	fn("blah")
}
