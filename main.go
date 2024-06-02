package main

import (
	"ascii/args"
	"ascii/output"
	"os"
)

func main() {

	draws, outputFile := args.Parse(os.Args[1:])
	if draws == nil {
		// nothing to draw
		output.PrintUsage()
	}
	output.Draw(draws, outputFile)

}
