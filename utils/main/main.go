package main

import (
	"fmt"

	"github.com/rayjonesjay/go/utils"
)

func main() {
	fd, e := utils.Opener("input.txt")
	if e != nil {
		fmt.Println(e)
		return
	}
	for {
		str, byt := utils.Reader(fd)
		if str == "" {
			if byt == 1 {
				break
			}

		}
		fmt.Println(str)

	}
}
