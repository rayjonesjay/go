package output

import (
	"fmt"
	"os"
	"strings"

	"ascii/args"
	"ascii/graphics"
	"ascii/sound"
	"ascii/special"
)

// Given a series of [args.DrawInfo] items, extract the drawing information and generate the expected graphics
func Draw(all []args.DrawInfo, outputFile string) {
	// The text received from the commandline may include special ASCII escape characters as \t, \a, \r, \v, \b, and \f
	// we handle such characters using the utilities from the `special` chars package
	hasBell := false
	for i := 0; i < len(all); i++ {
		d := all[i]
		if d.Text == "" {
			continue
		}
		// Handle the special characters \t, \b, \r, \f, \v
		// Interpret \t characters as two spaces
		d.Text = strings.ReplaceAll(d.Text, "\t", "  ")
		// The implementation of the `special` chars package didn't use the actual special characters; e.g.
		// the implementation used (\\r) instead of (\r)
		// functions in the special package only expect a single line of text for modification,
		// but our text may include multiple lines; thus, we feed each line separately to the functions
		d.Text = applyPerLine(d.Text, special.SlashB, "\b", "\\b")
		d.Text = applyPerLine(d.Text, special.SlashR, "\r", "\\r")
		d.Text = applyPerLine(d.Text, special.SlashV, "\v", "\\v")
		d.Text = applyPerLine(d.Text, special.SlashF, "\f", "\\f")

		// Handle \a
		if strings.ContainsRune(d.Text, '\a') {
			hasBell = true
			d.Text = strings.ReplaceAll(d.Text, "\a", "")
		}

		all[i] = d
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
			fmt.Fprintf(os.Stderr, "error writing  to  file\n")
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
