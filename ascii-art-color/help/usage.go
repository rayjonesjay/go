package help

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

var (
	UsageMessage = `Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"`

	hashOfUsageTxt = "c617203ce08ccc11b10aac61ddef73f4342419272253c87d51213356c9ca00e4"
)

// PrintUsage prints a short program usage to the standard output, then exits the program with a non-zero return code
func PrintUsage() {
	printUsage("plain/usage.txt", 0)
}

// PrintLongUsage prints the full program usage to the standard output,
// then exits the program with a non-zero return code
func PrintLongUsage() {
	printUsage("plain/usage.txt", -1)
	printUsage("plain/usage_long.txt", 0)
}

// printUsage prints the program usage as defined in the given file to the standard output,
// then exits the program with the given return code
func printUsage(fileName string, exitCode int) {

	usage, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred while reading file %v", err)
		os.Exit(1)
	}
	currentHash := sha256.Sum256(usage)
	currentHashString := hex.EncodeToString(currentHash[:])
	if currentHashString != hashOfUsageTxt {
		fmt.Fprintln(os.Stdout, UsageMessage)
		os.Exit(0)
	}

	fmt.Print(string(usage))
	if exitCode >= 0 {
		os.Exit(exitCode)
	}
}
