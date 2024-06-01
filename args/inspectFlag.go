package args

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// checks if the flag passed is valid --output=file.txt
func InspectFlagAndFile(flagAndFile string) (string, error) {
	if flagAndFile == "" {
		return "", fmt.Errorf("no flag passed")
	}

	hasWhiteSpace := strings.ContainsAny(flagAndFile, " \t\n\r")

	if hasWhiteSpace {
		return "", fmt.Errorf("wrong flag passed. might contain whitespaces. %q", flagAndFile)
	}

	flagPattern := `^(--output=)(\w{1,255}\.txt)`
	compiledPattern := regexp.MustCompile(flagPattern)

	resultOfMatch := compiledPattern.FindStringSubmatch(flagAndFile)
	validFlagFormat := "--output="

	if validFlagFormat != resultOfMatch[1] {
		return "", fmt.Errorf("wrong flag passed: expected %s got %s", validFlagFormat, resultOfMatch[1])
	}

	var fileToReceiveGraphics string
	if compiledPattern.MatchString(flagAndFile) {
		fileToReceiveGraphics = resultOfMatch[2]
	} else {
		index := strings.Index(flagAndFile, "=")
		if index != -1 {
			return "", fmt.Errorf("wrong file type passed: expected a file with '.txt' extension got %s", flagAndFile[index+1:])
		} else {
			return "", fmt.Errorf("wrong flag passed")
		}
	}

	return fileToReceiveGraphics, nil

}

func IsValidFlag(args []string) bool {

	if len(args) == 0 {
		return false
	}

	flag := args[0]

	isValidFlag, matchStringError := regexp.MatchString(`--output=\w{1,255}\.txt`, flag)

	if matchStringError != nil {
		fmt.Fprintf(os.Stdout, "Error with regexp matching string.")
		return false
	}

	return isValidFlag
}
