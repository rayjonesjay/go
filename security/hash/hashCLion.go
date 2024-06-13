package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

// HashFile calculates the sum of CLion-2024.1.3.tar.gz file
func  HashFile() {
	fullPath := "$HOME/Downloads/CLion-2024.1.3.tar.gz"
	fullPath = os.ExpandEnv(fullPath)
	dataBytes , err := os.ReadFile(fullPath)

	if err != nil {
		fmt.Println("error reading file:",err)
		os.Exit(1)
	}

	hash := sha256.Sum256(dataBytes)
	calculatedHash := hex.EncodeToString(hash[:])
	validSum := ("96f90b3898ce1393381f4ba4a46356b07993bb44e09680df383898bca3f508a3")
	if calculatedHash == validSum {
		fmt.Println("Sum is correct :)")
	}else{
		fmt.Println("Invalid Sum:")	
	}
}	


