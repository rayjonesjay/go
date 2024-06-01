package main

import (
	"ascii/args"
	"ascii/help"
	"ascii/output"
	"os"
)

func main() {
	mArgs := os.Args[1:]
	if len(mArgs) == 0 {
		return
	}

	draws, outputFile := args.Parse(mArgs)
	if draws == nil {
		// nothing to draw
		help.PrintUsage()
	}
	output.Draw(draws, outputFile)
}
