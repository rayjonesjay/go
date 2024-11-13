package output

import (
	"fmt"
	"os"
	"strings"

	"ascii/fmtx"

	"ascii/data"
	"ascii/graphics"
	"ascii/sound"
	"ascii/special"
)

// Draw given a [DrawInfo] item, extract the drawing information and generate the expected graphics
// The text received from the commandline may include special ASCII escape characters as \t, \a, \r, \v, \b, and \f
// we handle such characters using the utilities from the `special` chars package
func Draw(draws data.DrawInfo, outputFile string) {
	var hasBell bool
	draws.Text, hasBell = HandleSpecial(draws.Text)
	out, err := graphics.Draw(draws)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if outputFile != "" {
		fd, openError := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o666)
		if openError != nil || fd == nil {
			fmtx.FatalErrorf("error opening file\n")
			return
		}

		_, writeError := fd.WriteString(out)
		if writeError != nil {
			fmtx.FatalErrorf("error writing to file\n")
		}

		return
	}
	fmt.Print(out)
	if hasBell {
		// Some ASCII bell character was specified, play a beep sound
		sound.Beep()
	}
}

// HandleSpecial interprets the special characters \t, \b, \r, \f, \v, respectively
func HandleSpecial(text string) (string, bool) {
	hasBell := false
	if text != "" {
		// Handle the special characters \t, \b, \r, \f, \v, respectively
		text = strings.ReplaceAll(text, "\t", "  ")
		text = callFuncPerLine(text, special.EscapeB)
		text = callFuncPerLine(text, special.EscapeR)
		text = callFuncPerLine(text, special.EscapeF)
		text = callFuncPerLine(text, special.EscapeV)

		// Handle \a
		if strings.ContainsRune(text, '\a') {
			hasBell = true
			text = strings.ReplaceAll(text, "\a", "")
		}
	}

	return text, hasBell
}

// callFuncPerLine calls the function f separately to each line of the string s,
// and returns the joined results as a string
func callFuncPerLine(s string, f func(string) string) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		lines[i] = f(l)
	}
	return strings.Join(lines, "\n")
}
