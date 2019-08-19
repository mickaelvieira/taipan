package nonce

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

// returns a random integer between 2 values
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// returns the byte representation of a random minuscule, majuscule, number or special character
func randomSymbol() byte {
	return byte(randomInt(33, 126))
}

// Generate generates a random of the length provided
func Generate(len int) string {
	// https://golang.org/pkg/math/rand/#Rand.Seed
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = randomSymbol()
	}

	// we use sha256 algorithm here to get:
	// - a fix length string
	// - a URL safe string
	buf := sha256.New()
	buf.Write(bytes)
	b := buf.Sum(nil)

	return hex.EncodeToString(b)
}
