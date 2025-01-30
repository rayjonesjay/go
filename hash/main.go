package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	// dummy password
	password := []byte("supersecretpassword")
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Println("hashed password", string(hashedPassword))
	fmt.Println("******")
	err = cmp([]byte("$2a$10$9ZckUw7.nWKRJKeLR4PslOSqKCD.AsJ2sOLmt1ZoKfQj0Z.M.qH9e"), []byte("supersecretpassword"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("ok")
	}
}

func cmp(a, b []byte) error {
	return bcrypt.CompareHashAndPassword(a, b)
}
