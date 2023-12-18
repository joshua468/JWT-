package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateRandomKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func main() {
	// Specify the desired length of the secret key
	keyLength := 32

	// Generate a random secret key
	secretKey, err := generateRandomKey(keyLength)
	if err != nil {
		fmt.Println("Error generating secret key:", err)
		return
	}

	fmt.Println("Generated Secret Key:", secretKey)
}
