package args

import (
	"ascii/args/flags"
	. "ascii/data"
	"ascii/help"
	"fmt"
)

const (
	Shadow     = "shadow"
	Standard   = "standard"
	Thinkertoy = "thinkertoy"
)

// Parse takes the flag '--output=file.txt' together with text and style to be printed
func Parse(args []string) ([]DrawInfo, string) {
	lengthOfArguments := len(args)
	outputFile := ""

	// check if flag was passed and is valid
	if flags.IsValidFlag(args) {
		flagAndFile := args[0]
		var inspectError error
		outputFile, inspectError = flags.InspectFlagAndFile(flagAndFile)
		if inspectError != nil {
			fmt.Printf("Usage Error: %s\n", inspectError.Error())
			help.PrintUsage()
		}
		args = args[1:]
		lengthOfArguments = lengthOfArguments - 1
	}

	if lengthOfArguments < 1 {
		return nil, outputFile
	} else if lengthOfArguments == 1 {
		text := args[0]
		drawInfo := DrawInfo{Text: Escape(text), Style: Standard}
		return []DrawInfo{drawInfo}, outputFile
	} else if lengthOfArguments == 2 {
		text, style := args[0], args[1]
		drawInfo := DrawInfo{Text: Escape(text), Style: style}
		return []DrawInfo{drawInfo}, outputFile
	}

	return nil, outputFile
}
