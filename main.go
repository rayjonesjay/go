package main

import (
	"fmt"
	"os"

	"ascii/args"
	"ascii/graphics"
)

func main() {
	draws := args.Parse(os.Args[1:])
	if draws == nil {
		// nothing to draw, usage error
		printUsage()
	}
	m := graphics.ReadBanner("banners/standard.txt")
	var result [][]string
	for _, di := range draws {
		text := di.Text
		for _, char := range text {
			result = append(result, m[char])
		}
	}

	for i := 0; i < 8; i++ {
		for _, element := range result {
			fmt.Print(element[i])
		}
		fmt.Println()
	}

	// for _, element := range result {
	// 	for _, value := range element {
	// 		fmt.Println(value)
	// 	}
	// }

	// draw(draws)
}

// Prints the program usage to the standard output, then exits the program with a non-zero return code
func printUsage() {
	usage, err := os.ReadFile("plain/usage.txt")
	if err != nil {
		// We couldn't read the usage text our program was shipped with!
		_, _ = fmt.Fprintln(os.Stderr, "Improper program installation. Re-installation recommended!!")
		os.Exit(5)
	}
	fmt.Print(string(usage))
	os.Exit(1)
}

// Given a series of [DrawInfo] items, extract the drawing information and generate the expected graphics
func draw(all []args.DrawInfo) {
	// TODO
}
