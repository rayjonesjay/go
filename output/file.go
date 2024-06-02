package output

import (
	"fmt"
	"os"
	"strings"

	. "ascii/data"
	"ascii/graphics"
	"ascii/sound"
	"ascii/special"
)

// Draw given a series of [DrawInfo] items, extract the drawing information and generate the expected graphics
func Draw(all []DrawInfo, outputFile string) {
	// The text received from the commandline may include special ASCII escape characters as \t, \a, \r, \v, \b, and \f
	// we handle such characters using the utilities from the `special` chars package
	hasBell := false
	for i := 0; i < len(all); i++ {
		drawInfo := all[i]
		if drawInfo.Text == "" {
			continue
		}
		// Handle the special characters \t, \b, \r, \f, \v
		drawInfo.Text = strings.ReplaceAll(drawInfo.Text, "\t", "  ")
		drawInfo.Text = applyPerLine(drawInfo.Text, special.SlashB, "\b", "\\b")
		drawInfo.Text = applyPerLine(drawInfo.Text, special.SlashR, "\r", "\\r")
		drawInfo.Text = applyPerLine(drawInfo.Text, special.SlashV, "\v", "\\v")
		drawInfo.Text = applyPerLine(drawInfo.Text, special.SlashF, "\f", "\\f")

		// Handle \a
		if strings.ContainsRune(drawInfo.Text, '\a') {
			hasBell = true
			drawInfo.Text = strings.ReplaceAll(drawInfo.Text, "\a", "")
		}

		all[i] = drawInfo
	}

	out := graphics.Draw(all)

	if outputFile != "" {
		fd, openError := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o666)
		if openError != nil {
			fmt.Fprintf(os.Stderr, "error opening file\n")
			os.Exit(1)
		}

		_, writeError := fd.WriteString(out)
		if writeError != nil {
			fmt.Fprintf(os.Stderr, "error writing to file\n")
			os.Exit(1)
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
