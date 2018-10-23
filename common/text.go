package common

import "crypto/rand"

// GenerateRandomText - Generate random text
func GenerateRandomText() (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err == nil {
		for i, r := range b {
			b[i] = letters[r%byte(len(letters))]
		}
		return string(b), err
	}

	return "", err
}
