package fmtx

import (
	"ascii/colors"
	"ascii/terminal"
	"fmt"
	"os"
)

// Errorf writes the formated output string to the console standard error stream
func Errorf(format string, args ...interface{}) {
	// We are writing to the standard err, assumed to be the console, ignore any errors that may arise
	if terminal.IsTerminal() {
		format = colors.Red + format + colors.Reset
	}
	_, err := fmt.Fprintf(os.Stderr, format, args...)
	if err != nil {
		fmt.Println("Couldn't write to Stderr:", err)
		os.Exit(1)
	}
}

// FatalErrorf writes the formated output string to the console standard error stream,
// then exits the program unsuccessfully
func FatalErrorf(format string, args ...interface{}) {
	Errorf(format, args...)
	os.Exit(1)
}
