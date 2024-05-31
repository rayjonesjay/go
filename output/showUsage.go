package output

import (
	"fmt"
	"os"
)

// Prints the program usage to the standard output, then exits the program with a non-zero return code
func PrintUsage() {
	usage, err := os.ReadFile("plain/usage.txt")
	if err != nil {
		// We couldn't read the usage text our program was shipped with!
		_, _ = fmt.Fprintln(os.Stderr, "Improper program installation. Re-installation recommended!!")
		os.Exit(1)
	}
	fmt.Print(string(usage))
	os.Exit(1)
}
