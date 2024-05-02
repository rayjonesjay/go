package main

import (
	"ascii/graphics"
	"ascii/special"
	"fmt"
	"os"
	"strings"

	"ascii/args"
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
		// FIXME:
		// current implementation of the special chars package didn't use the actual special characters; e.g.
		//the implementation used (\\r) instead of (\r)
		d.Text = strings.ReplaceAll(d.Text, "\b", "\\b")
		d.Text = strings.ReplaceAll(d.Text, "\r", "\\r")

		// Handle the special characters \b, \r, \0
		d.Text = special.SlashB(d.Text)
		//d.Text = special.SlashR(d.Text) // This fails: 'Go\nHello\r12ere'? this too: 'Go'?

		all[i] = d
	}
	out := graphics.Draw(all)
	fmt.Print(out)
}
