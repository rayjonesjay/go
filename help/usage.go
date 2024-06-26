package help

import (
	"fmt"
	"os"
)

// PrintUsage prints a short program usage to the standard output, then exits the program with a non-zero return code
func PrintUsage() {
	printUsage("plain/usage.txt", 1)
}

// PrintLongUsage prints the full program usage to the standard output,
// then exits the program with a non-zero return code
func PrintLongUsage() {
	printUsage("plain/usage.txt", -1)
	printUsage("plain/usage_long.txt", 1)
}

// printUsage prints the program usage as defined in the given file to the standard output,
// then exits the program with the given return code
func printUsage(fileName string, exitCode int) {
	usage, err := os.ReadFile(fileName)
	if err != nil {
		// We couldn't read the usage text our program was shipped with!
		_, _ = fmt.Fprintln(os.Stderr, "Improper program installation. Re-installation recommended!!")
		os.Exit(1)
	}
	fmt.Print(string(usage))
	if exitCode >= 0 {
		os.Exit(exitCode)
	}
}
