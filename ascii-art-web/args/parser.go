package args

import (
	"ascii/args/flags"
	"ascii/data"
	"ascii/help"
	"regexp"
	"slices"
)

const (
	Standard = "standard"
)

// ParserOut structures the program arguments for simpler access to individual parsed arguments
type ParserOut struct {
	Draws      *data.DrawInfo
	OutputFile string
}

// Parse extracts flags and positional arguments from the given command-line arguments, then interprets the expected
// text to be drawn and with which banner style to be used.
// Implementation Details:
// - This also extracts the output file for the graphics if specified.
// - In case the --help flag is specified, then it displays the program help and exits
func Parse(args []string) ParserOut {
	mFlags, mArgs := parseFlags(args)

	// If a user specified the --help flag, then we only need to display the program help and exit
	if slices.Contains(mFlags, "--help") {
		help.PrintLongUsage()
	}

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

// parseFlags extracts flags and positional arguments from the given command-line arguments
func parseFlags(args []string) ([]string, []string) {
	var mFlags []string
	re := regexp.MustCompile(`^--.*`)

	argStart := 0
	for i, mArg := range args {
		if re.MatchString(mArg) {
			mFlags = append(mFlags, mArg)
			argStart = i + 1
		}
	}

	return mFlags, args[argStart:]
}
