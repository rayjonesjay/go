package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println("hello world")
	fmt.Println(Fooer(13))
}

func Fooer(input int) string {

	isFoo := (input % 3) == 0
	if isFoo {
		return "foo"
	}
	return strconv.Itoa(input)
}
