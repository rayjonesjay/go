package main

import (
	"ascii/args"
	"ascii/graphics"
	"ascii/special"
	"fmt"
	"os"
	"strings"
)

func main() {
	draws := args.Parse(os.Args[1:])
	if draws == nil {
		// nothing to draw, usage error
		printUsage()
	}
	draw(draws)
}

// Prints the program usage to the standard output, then exits the program with a non-zero return code
func printUsage() {
	usage, err := os.ReadFile("plain/usage.txt")
	if err != nil {
		// We couldn't read the usage text our program was shipped with!
		_, _ = fmt.Fprintln(os.Stderr, "Improper program installation. Re-installation recommended!!")
		os.Exit(1)
	}
	fmt.Print(string(usage))
	os.Exit(1)
}

// Given a series of [args.DrawInfo] items, extract the drawing information and generate the expected graphics
func draw(all []args.DrawInfo) {
	// The text received from the commandline may include special ASCII escape characters as \r, \v, \b
	//we handle such using the utilities from the special chars package
	for i := 0; i < len(all); i++ {
		d := all[i]
		if d.Text == "" {
			continue
		}
		// Handle the special characters \t, \b, \r, \f, \v
		// Interpret \t characters as two spaces
		d.Text = strings.ReplaceAll(d.Text, "\t", "  ")
		// The implementation of the `special` chars package didn't use the actual special characters; e.g.
		//the implementation used (\\r) instead of (\r)
		// functions in the special package only expect a single line of text for modification,
		//but our text may include multiple lines; thus, we feed each line separately to the functions
		d.Text = applyPerLine(d.Text, special.SlashB, "\b", "\\b")
		d.Text = applyPerLine(d.Text, special.SlashR, "\r", "\\r")
		d.Text = applyPerLine(d.Text, special.SlashV, "\v", "\\v")
		d.Text = applyPerLine(d.Text, special.SlashF, "\f", "\\f")

		all[i] = d
	}
	out := graphics.Draw(all)
	fmt.Print(out)
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
