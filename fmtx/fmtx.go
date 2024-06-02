package fmtx

import (
	"fmt"
	"os"
)

// Errorf writes the formated output string to the console standard error stream
func Errorf(format string, args ...interface{}) {
	// We are writing to the standard err, assumed to be the console, ignore any errors that may arise
	_, err := fmt.Fprintf(os.Stderr, format, args)
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
