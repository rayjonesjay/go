package args

import (
	"ascii/args/flags"
	"ascii/data"
)

const (
	Shadow     = "shadow"
	Standard   = "standard"
	Thinkertoy = "thinkertoy"
)

// ParserOut structures the program arguments for simpler access to individual parsed arguments
type ParserOut struct {
	Draws      *data.DrawInfo
	OutputFile string
}

// Parse takes the flag '--output=file.txt' together with text and style to be printed
func Parse(args []string) ParserOut {
	lengthOfArguments := len(args)
	outputFile := flags.InspectFlagAndFile(args)

	args = args[1:]
	lengthOfArguments = lengthOfArguments - 1

	if lengthOfArguments < 1 {
		return ParserOut{OutputFile: outputFile}
	} else if lengthOfArguments == 1 {
		text := args[0]
		drawInfo := data.DrawInfo{Text: Escape(text), Style: Standard}
		return ParserOut{Draws: &drawInfo, OutputFile: outputFile}
	} else if lengthOfArguments == 2 {
		text, style := args[0], args[1]
		drawInfo := data.DrawInfo{Text: Escape(text), Style: style}
		return ParserOut{Draws: &drawInfo, OutputFile: outputFile}
	}

	return ParserOut{OutputFile: outputFile}
}
