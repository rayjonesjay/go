package args

import (
	"ascii/args/flags"
	"ascii/data"
	"ascii/help"
	"regexp"
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
	mFlags, mArgs := parseFlags(args)

	var outputFile string
	if len(mFlags) > 0 {
		outputFile = flags.InspectFlagAndFile(mFlags)
	}

	args = mArgs
	lengthOfArguments := len(mArgs)

	if lengthOfArguments < 1 {
		return ParserOut{OutputFile: outputFile}
	} else if lengthOfArguments == 1 {
		text := mArgs[0]
		drawInfo := data.DrawInfo{Text: Escape(text), Style: Standard}
		return ParserOut{Draws: &drawInfo, OutputFile: outputFile}
	} else if lengthOfArguments == 2 {
		text, style := mArgs[0], mArgs[1]
		drawInfo := data.DrawInfo{Text: Escape(text), Style: style}
		return ParserOut{Draws: &drawInfo, OutputFile: outputFile}
	} else {
		help.PrintUsage()
	}

	return ParserOut{OutputFile: outputFile}
}

func parseFlags(args []string) ([]string, []string) {
	var f []string
	re := regexp.MustCompile(`^--.*`)

	argStart := 0
	for i, s := range args {
		if re.MatchString(s) {
			f = append(f, s)
			argStart = i + 1
		}
	}

	return f, args[argStart:]
}
