package output

import (
	"ascii/fmtx"
	"fmt"
	"os"
	"strings"

	"ascii/data"
	"ascii/graphics"
	"ascii/sound"
	"ascii/special"
)

// Draw given a [DrawInfo] item, extract the drawing information and generate the expected graphics
// The text received from the commandline may include special ASCII escape characters as \t, \a, \r, \v, \b, and \f
// we handle such characters using the utilities from the `special` chars package
func Draw(draws data.DrawInfo, outputFile string) {
	hasBell := false
	if draws.Text != "" {
		// Handle the special characters \t, \b, \r, \f, \v, respectively
		draws.Text = strings.ReplaceAll(draws.Text, "\t", "  ")
		draws.Text = callFuncPerLine(draws.Text, special.EscapeB)
		draws.Text = callFuncPerLine(draws.Text, special.EscapeR)
		draws.Text = callFuncPerLine(draws.Text, special.EscapeF)
		draws.Text = callFuncPerLine(draws.Text, special.EscapeV)

		// Handle \a
		if strings.ContainsRune(draws.Text, '\a') {
			hasBell = true
			draws.Text = strings.ReplaceAll(draws.Text, "\a", "")
		}
	}

	out := graphics.Draw(draws)

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

// callFuncPerLine calls the function f separately to each line of the string s,
// and returns the joined results as a string
func callFuncPerLine(s string, f func(string) string) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		lines[i] = f(l)
	}
	return strings.Join(lines, "\n")
}
