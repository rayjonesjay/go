package security

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

var hashMap = map[string]string{
	"standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
}

// Calculate the hash of the current file to ensure file's integrity.
func GetDataAndCalculateHash(filePath, fileName string) bool {
	fileDataInBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading banner file: %q\n", filePath)
		os.Exit(1)
	}

	validHash, ok := hashMap[fileName]
	if !ok {
		return true
	}

	sha265Hash := sha256.Sum256(fileDataInBytes)
	calculatedHash := hex.EncodeToString(sha265Hash[:])
	return calculatedHash == validHash
}
