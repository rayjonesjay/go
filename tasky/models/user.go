package models

import (
	"context"
	"fmt"
	"log"
	"tasky/db"
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateUser(email, passwordHash string) error {
	query := `
	INSERT INTO users (email,password_hash)
	VALUES ($1, $2)
	`
	_, err := db.Conn.Exec(context.Background(), query, email, passwordHash)
	if err != nil {
		log.Printf("error creating user: %v", err)
		return err
	}

	fmt.Println("user created successfully")
	return nil
}

func GetUserByEmail(email string) (*User, error) {
	query := `
	SELECT id, email, password_hash, created_at
	FROM users
	WHERE email = $1`
	var u User

	err := db.Conn.QueryRow(context.Background(), query, email).Scan(
		&u.Id,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
	)
	if err != nil {
		log.Printf("error retrieving user: %v", err)
		return nil, err
	}
	return &u, nil
}
