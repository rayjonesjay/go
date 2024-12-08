package main

import "fmt"

func main() {
	fmt.Println(a())

	fmt.Println(b(funcs{add, mul, div}))
}

func x(a int, b int) int {
	return a + b
}

func add(a int) int {
	return a + a
}

func mul(a int) int {
	return a * a
}

func div(a int) int {
	return a / a
}

/*
a slice of functions that take an integer and return an int
*/
type funcs []func(int) int

func b(fs funcs) int {
	n := 1
	for _, f := range fs {
		n += f(n)
	}
	return n
}
func a(funcs ...func(int) int) int {
	var n = 1
	for _, f := range funcs {
		fmt.Println("n is now", n)
		n += f(n)
	}
	return n
}
