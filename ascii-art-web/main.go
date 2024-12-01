package main

import (
	"os"

	"ascii/args"
	"ascii/help"
	"ascii/output"
	"ascii/web"
)

func main() {
	mArgs := os.Args[1:]
	if len(mArgs) == 0 {
		web.Server()
		return
	}

	parse := args.Parse(mArgs)
	if parse.Draws == nil {
		// nothing to draw
		help.PrintUsage()
		return
	}
	output.Draw(*parse.Draws, parse.OutputFile)
}
