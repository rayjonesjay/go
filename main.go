package main

import (
	"ascii/args"
	"ascii/help"
	"ascii/output"
	"os"
)

func main() {

	draws, outputFile := args.Parse(os.Args[1:])
	if draws == nil {
		// nothing to draw
		help.PrintUsage()
	}
	output.Draw(draws, outputFile)

}
