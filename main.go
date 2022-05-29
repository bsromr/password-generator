package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)

	return fmt.Sprintf("%x", b)[:length]
}

func main() {
	fmt.Println(generatePassword(12))
}
