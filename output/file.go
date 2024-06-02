package output

import (
	"ascii/fmtx"
	"fmt"
	"os"
	"strings"

	. "ascii/data"
	"ascii/graphics"
	"ascii/sound"
	"ascii/special"
)

// Draw given a [DrawInfo] item, extract the drawing information and generate the expected graphics
// The text received from the commandline may include special ASCII escape characters as \t, \a, \r, \v, \b, and \f
// we handle such characters using the utilities from the `special` chars package
func Draw(draws DrawInfo, outputFile string) {
	hasBell := false
	if draws.Text != "" {
		// Handle the special characters \t, \b, \r, \f, \v
		draws.Text = strings.ReplaceAll(draws.Text, "\t", "  ")
		draws.Text = applyPerLine(draws.Text, special.SlashB, "\b", "\\b")
		draws.Text = applyPerLine(draws.Text, special.SlashR, "\r", "\\r")
		draws.Text = applyPerLine(draws.Text, special.SlashV, "\v", "\\v")
		draws.Text = applyPerLine(draws.Text, special.SlashF, "\f", "\\f")

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

// applyPerLine applies the function f separately to each line of the string s, and returns the results as a string
func applyPerLine(s string, f func(string) string, real, escape string) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		sections := strings.Split(l, escape)
		for j, sn := range sections {
			sections[j] = f(strings.ReplaceAll(sn, real, escape))
		}
		lines[i] = strings.Join(sections, escape)
	}
	return strings.Join(lines, "\n")
}
