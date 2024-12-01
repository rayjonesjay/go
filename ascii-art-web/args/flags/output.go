package flags

import (
	"ascii/fmtx"
	"ascii/help"
	"regexp"
	"strings"
	"unicode"
)

// InspectFlagAndFile checks if the flag passed is valid --output=file.txt
func InspectFlagAndFile(args []string) string {
	if len(args) > 1 {
		help.PrintUsage()
	}

	// --output=<file.txt>
	flagAndFile := args[0]
	for _, r := range flagAndFile {
		if !unicode.IsPrint(r) {
			help.PrintUsage()
		}
	}

	// go run . --output=. Hello standard
	compiledFlagPattern := regexp.MustCompile(`^(--)(output=)(.+)$`)
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
