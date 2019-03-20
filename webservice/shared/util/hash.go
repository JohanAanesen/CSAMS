package util

import "golang.org/x/crypto/bcrypt"

// GenerateFromPassword returns hashed string with cost 14
func GenerateFromPassword(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	return string(bytes), err
}

// CompareHashAndPassword compares input with hash, returns error if any
func CompareHashAndPassword(input, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))
}