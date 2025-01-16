// package one_time_pad implements the one time pad encryption algorithm using the Go
// programming language.
package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

// randomness is vital in cryptography
// genKey is used to generate a random key
func genKey(length_of_key int) ([]byte, error) {
	key := make([]byte, length_of_key)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// a function to encrypt the message
func encrypt(message, key []byte) (cipherText []byte, _ error) {
	// the length of the message is supposed to be at least same length as key
	if len(message) != len(key) {
		return nil, fmt.Errorf("message and key must have the same length")
	}

	// cipher text is the text that has been encrypted with the key
	cipherText = make([]byte, len(message))

	for i := 0; i < len(message); i++ {
		// each bit of the message is XORED with a corresponding bit in the key value
		cipherText[i] = message[i] ^ key[i]
	}
	return cipherText, nil
}

func decrypt(ciperText, key []byte) (plainText []byte, _ error) {
	if len(ciperText) != len(key) {
		return nil, errors.New("ciphertext and key must have the same length")
	}

	plainText = make([]byte, len(ciperText))
	for i := 0; i < len(ciperText); i++ {
		plainText[i] = ciperText[i] ^ key[i]
	}

	return plainText, nil
}

func main() {
	message := "Some put it on the devil when they fall short, I put it on my ego- the lord of all lords"
	key, _ := genKey(len(message))

	cipherText, _ := encrypt([]byte(message), key)
	fmt.Printf("CIPHER: %x\n", cipherText)

	// utf := hex.EncodeToString(plainText)
	// utft, _ := hex.DecodeString(utf)
	// fmt.Println("???", string(utft))
	// fmt.Println("utf  ", utf)
	println()
	plainText, _ := decrypt([]byte(cipherText), key)
	fmt.Printf("PLAINTEXT: %s\n", string(plainText))
}
