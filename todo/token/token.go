package token

import (
	"fmt"
	"time"

	// "github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte(LoadEnv("SECRET"))

// CreateToken generates a JWT token with claims
func CreateToken(username string) (string, error) {
	// create a new jwt token with claims
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": username,                           // subject (user identifier)
			"iss": "todoApp",                          // Issuer
			"aud": getRole(username),                  // Audience (user role)
			"exp": time.Now().Add(time.Minute).Unix(), // expiration time
			"iat": time.Now().Unix(),                  // Issuead at
		})
	fmt.Printf("token with claims %+v\n", claims)
	tokenString, err := claims.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// This simple function checks if the username is senior and returns the role
// senior, otherwise, it defaults to employee. It will be used in the CreateToken function
// to set the audience claim of the JWT
func getRole(username string) string {
	if username == "senior" {
		return "senior"
	}
	return "employee"
}
