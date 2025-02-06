package token

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	EnvMap           = make(map[string]string)
	PathToEnvFile, _ = filepath.Abs(".env")
	Tok              string
)

func init() {
	fmt.Println("initializing from token package...")
	// read from .env file in the project root
	data, err := os.ReadFile(PathToEnvFile)
	if err != nil {
		fmt.Println("environment variables initialization failed", err)
		os.Exit(2)
	}
	Tok = string(data)
	splitAtNewLine := strings.Split(Tok, "\n")
	for _, line := range splitAtNewLine {
		keyValue := strings.Split(line, "=")
		EnvMap[keyValue[0]] = keyValue[1]
	}
	fmt.Println("done...")
	fmt.Println(EnvMap)
}

// LoadEnv returns the specified value of the environment variable based on the key
func LoadEnv(key string) (value string) {
	// check if key exists
	value, ok := EnvMap[key]
	if !ok {
		return ""
	}
	return value
}
