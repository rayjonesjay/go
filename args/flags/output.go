package flags

import (
	"ascii/fmtx"
	"ascii/help"
	"regexp"
	"strings"
)

// InspectFlagAndFile checks if the flag passed is valid --output=file.txt
func InspectFlagAndFile(args []string) string {
	if len(args) > 3 {
		help.PrintUsage()
	}

	// --output=<file.txt>
	flagAndFile := args[0]

	// remove trailing and leading spaces
	flagAndFile = strings.TrimSpace(flagAndFile)

	if strings.Contains(flagAndFile, "\n\r\t") {
		help.PrintUsage()
	}
	// go run . --output=. Hello standard

	flagPattern := `^(--)(output=)([^\s]+)$`

	compiledFlagPattern := regexp.MustCompile(flagPattern)

	matches := compiledFlagPattern.FindStringSubmatch(flagAndFile)

	if matches == nil {
		help.PrintUsage()
	}

	// if the flag does not contain --output= and does not start with --
	if !(strings.HasPrefix(flagAndFile, matches[1]) && strings.Contains(flagAndFile, matches[2])) {
		help.PrintUsage()
	}

	if matches[3] == "." || matches[3] == ".." {
		fmtx.FatalErrorf("invalid output file: %q\n", matches[3])
	}
	return matches[3]
}
