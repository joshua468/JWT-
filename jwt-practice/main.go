package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Secret key (this should be kept secret and not shared)
var secretKey = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MDI5MjEwMTUsIm5hbWUiOiJKb2huIERvZSIsInN1YiI6IjEyMzQ1Njc4OTAifQ.QprXpFKpwajgonTl10-7LHJxFTQJS4fC4fMGQj6GhCA")

func main() {
	// Example payload with user information
	payload := jwt.MapClaims{
		"sub":  "1234567890",
		"name": "John Doe",
		"iat":  time.Now().Unix(),
	}

	// Create a JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error creating JWT:", err)
		return
	}

	fmt.Println("Generated JWT:")
	fmt.Println(signedToken)
	fmt.Println()

	// Verify the JWT
	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("Error parsing JWT:", err)
		return
	}

	if parsedToken.Valid {
		fmt.Println("Decoded Payload:")
		fmt.Println(parsedToken.Claims)
	} else {
		fmt.Println("Invalid token.")
	}
}
