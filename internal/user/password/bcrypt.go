package password

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	const LEN = 8
	if len(password) < LEN {
		return "", errors.New(fmt.Sprintf("should be %d letters long", LEN))
	}

	if !hasLowerCase(password) {
		return "", errors.New("should contain lowercase letters")
	}

	if !hasUpperCase(password) {
		return "", errors.New("should contain uppercase letters")
	}

	if !hasSpecialChar(password) {
		return "", errors.New("should contain special characters")
	}

	if !hasDigit(password) {
		return "", errors.New("should contain digits")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func IsPasswordMatch(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
