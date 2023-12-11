package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	key := make([]byte, 32) // 32 bytes will result in a 256-bit key
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}
	encodedKey := base64.URLEncoding.EncodeToString(key)
	fmt.Println("Generated Key:", encodedKey)
}
