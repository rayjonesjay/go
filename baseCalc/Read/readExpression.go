// Reads the expression and ensures it is in the correct format before evaluation.
package Read

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"baseCalc/Evaluate"
	"regexp"
)
type Mode struct {
	To int 
	From int 
}
var CurrentMode = Mode{From:2,To:10}

func ReadExpression() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

	for {
		writer.WriteString("[in ]: ")
		writer.Flush()
		input, readError := reader.ReadString('\n')
		if readError != nil {
			fmt.Fprintf(os.Stdout, "read error")
		}
		// remove newlines, whitespaces ..etc
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "quit" || input == ".q" || input == "exit" || input == ".e" {
			writer.WriteString("exiting...\n")
			writer.Flush()
			os.Exit(0)
		}

		var output string = ""

		// check if the user is setting mode to different mode
		if IsMode(input) {
			CurrentMode = SetBase(input)
			writer.WriteString("[mode set]: " + fmt.Sprintf("%v -> %v\n", CurrentMode.From, CurrentMode.To))
			writer.Flush()
		} else if IsExpression(input) {
			output = Evaluate.EvaluateExpression(input, CurrentMode.From, CurrentMode.To)
			writer.WriteString("[out]: " + output + "\n")
			writer.Flush()
		} else {
			writer.WriteString("invalid base number ensure base is in range [2 to 10]\nUsage: [in]: mode=2->10\n")
			writer.Flush()
		}
	}
}


// Check if user input is mode and returns from(int) to(int)
func SetBase(input string) Mode {
	pattern := `^mode=(\d{1,2})(->|<-)(\d{1,2})$`

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 4 {
		fmt.Println("Invalid mode format.")
		return CurrentMode
	}

	left := matches[1]
	arrow := matches[2]
	right := matches[3]

	if IsDigitRange(left) && IsDigitRange(right) {
		from, to := arrowDirection(left, arrow, right)
		fromInt := stringToInt(from)
		toInt := stringToInt(to)
		return Mode{From: fromInt, To: toInt}
	}

	fmt.Println("Base out of range. Must be between 2 and 10.")
	return CurrentMode
}

func IsExpression(input string) bool {
	hexPattern := `^[0-9-a-f]+$`
	binaryPattern := `^[0-1]+$`
	decimalPattern := `^[0-9]+$`
	octalPattern := `^[0-7]+$`
	re := regexp.MustCompile(hexPattern)
	if re.MatchString(input){
		return true 
	}
	re = regexp.MustCompile(binaryPattern)
	if re.MatchString(input){
		return true 
	}
	re = regexp.MustCompile(octalPattern)
	if re.MatchString(input){
		return true 
	}
	re = regexp.MustCompile(decimalPattern)
	return re.MatchString(input)
}


func IsMode(input string) bool {
	pattern := regexp.MustCompile(`^mode=\d{1,2}(->|<-)\d{1,2}$`)
	return pattern.MatchString(input)
}



func stringToInt(s string) int {
	if s[0] == '-'{
		fmt.Fprintf(os.Stderr, "cannot compute base of negatives")
		os.Exit(1)
	}

	if s[0] == '+'{
		s = s[1:]
	}

	var result int = 0
	for _, char := range  s {
		result = result * 10 + int(char-'0')
	}

	return result
}


// arrow direction tells the program to convert from which base to which base.
func arrowDirection(left,arrow,right string) (from,to string) {
	if arrow == "->" {
		return left,right
	}
	return right,left
}

func IsDigitRange(s string) bool {
	intS := stringToInt(s)
	if intS >= 2 && intS <= 10{
		return true
	}
	return false
}