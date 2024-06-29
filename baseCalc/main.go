package main 

import (
	"bufio"
	"os"
	"baseCalc/Read"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	// welcome message for user
	writer.WriteString("welcome to base converter by https://github.com/rayjonesjay/\ntype \"quit or .q\" to exit the program.\ndefault mode=10->2\n")
	writer.WriteString("--------------------------------------------\n")
	writer.Flush()
	Read.ReadExpression()
}