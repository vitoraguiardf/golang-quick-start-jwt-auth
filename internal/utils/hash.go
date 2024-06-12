package utils

import "golang.org/x/crypto/bcrypt"

func HashCreate(password string) (string, error) {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), 14); err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

func HashCompare(hashed string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
