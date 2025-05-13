package password

import "unicode"

func hasUpperCase(password string) bool {
	for _, char := range password {
		if unicode.IsUpper(char) {
			return true
		}
	}

	return false
}

func hasLowerCase(password string) bool {
	for _, char := range password {
		if unicode.IsLower(char) {
			return true
		}
	}

	return false
}

func hasSpecialChar(password string) bool {
	for _, char := range password {
		if !unicode.IsDigit(char) && !unicode.IsLetter(char) {
			return true
		}
	}

	return false
}

func hasDigit(password string) bool {
	for _, char := range password {
		if unicode.IsDigit(char) {
			return true
		}
	}

	return false
}
