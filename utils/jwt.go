package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

func GenerateJWT(currentUserId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"currentUserId": currentUserId,
			"exp":           time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(getSecretKey())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// authenticationHeader sample value Authorization: Bearer <token>
func GetCurrentUserId(authenticationHeader string) (int, error) {
	tokenString, err := extractJWT(authenticationHeader)
	if err != nil {
		return 0, err
	}

	token, err := validate(tokenString)
	if err != nil {
		return 0, err
	}

	currentUserId, err := parseClaims(token)
	if err != nil {
		return 0, err
	}

	return currentUserId, nil
}

// Sample authenticationHeader value "Authorization: Bearer <token>"
// Token string return value sample eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30
func extractJWT(authenticationHeader string) (tokenString string, err error) {
	if strings.TrimSpace(authenticationHeader) == "" {
		return "", errors.New("authorization header is required")
	}

	parts := strings.SplitN(authenticationHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}

// Sample tokenString value eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30
func validate(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return getSecretKey(), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return token, nil
}

func parseClaims(token *jwt.Token) (currentUserId int, err error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}

	// JWT numbers are parsed as float64
	currentUserIdFloat, ok := claims["currentUserId"].(float64)
	if !ok {
		return 0, err
	}

	return int(currentUserIdFloat), nil
}

func getSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}
