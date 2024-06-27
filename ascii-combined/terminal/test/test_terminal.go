package main

import (
	"ascii/fmtx"
	"ascii/terminal"
	"fmt"
)

func leftAlign(text string) {
	fmt.Println(text)
}

func rightAlign(text string, width int) {
	padding := width - len(text)
	if padding > 0 {
		fmt.Printf("%*s\n", width, text)
	} else {
		fmt.Println(text)
	}
}

func main() {
	width, _, err := terminal.GetTerminalSize()
	if err != nil {
		fmtx.FatalErrorf("error getting terminal size\n%v\n", err)
		return
	}

	text := "Hello, World!"

	fmt.Println("Left Aligned:")
	leftAlign(text)

	fmt.Println("\nRight Aligned:")
	rightAlign(text, width)
}
