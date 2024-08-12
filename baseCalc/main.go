package main

import (
	"baseCalc/Read"
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	// welcome message for user
	someMessage := fmt.Sprintf("\ndefault mode=%d->%d\n", Read.CurrentMode.From, Read.CurrentMode.To)
	writer.WriteString("welcome to base converter by https://github.com/rayjonesjay/\ntype \"quit or .q\" to exit the program." + someMessage)
	writer.WriteString("--------------------------------------------\n")
	writer.Flush()
	Read.ReadExpression()
}
