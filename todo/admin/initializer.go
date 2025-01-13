package admin

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	API_KEY string
	API_URL string
)

func LoadEnv() {
	// read the .env file
	abs, _ := filepath.Abs(".env")
	env, err := os.Open(abs)
	if err != nil {
		log.Println("err", err)
		return
	}
	scanner := bufio.NewScanner(env)

	for scanner.Scan() {
		line := scanner.Text()
		keyValue := strings.Split(line, "=")
		os.Setenv(keyValue[0], keyValue[1])
	}
	fmt.Println("done loading environment variables")
	var isFound bool
	API_KEY, isFound = os.LookupEnv("API_KEY")
	if !isFound {
		fmt.Println("KEY NOT FOUND")
		return
	}
	API_URL, isFound = os.LookupEnv("API_URL")

	if !isFound {
		fmt.Println("URL NOT FOUND")
		return
	}
}

func init() {
	LoadEnv()
}
