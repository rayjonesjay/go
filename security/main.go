package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	fullPath := "$HOME/Downloads/CLion-2024.1.3.tar.gz"
	fullPath = os.ExpandEnv(fullPath)
	dataBytes , err := os.ReadFile(fullPath)

	if err != nil {
		fmt.Println("error reading file:",err)
		os.Exit(1)
	}

	hash := sha256.Sum256(dataBytes)
	hashString := hex.EncodeToString(hash[:])
	fmt.Println(hashString=="96f90b3898ce1393381f4ba4a46356b07993bb44e09680df383898bca3f508a3")
}


