package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"strings"
	"time"
)

// authenticationHeader = Authorization: Bearer <jwt>
// tokenString referred as the jwt but its not validated yet
// token referred as the jwt and its validated

func GenerateJWT(id int) (string, error) {
	expirationInMinute, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_IN_MINUTE"))
	if err != nil {
		return "", err
	}

	now := time.Now()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"iat": now,
		"exp": now.Add(time.Duration(expirationInMinute) * time.Minute).Unix(),
		"iss": os.Getenv("JWT_ISSUER"),
		"aud": os.Getenv("JWT_AUDIENCE"),
	}).SignedString(getSecretKey())

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetSubject(authenticationHeader string) (int, error) {
	sub, err := validateAndParse(authenticationHeader, "sub")
	if err != nil {
		return 0, err
	}

	id, ok := sub.(float64)
	if !ok {
		return 0, errors.New("id is not a number")
	}

	return int(id), nil
}

// authenticationHeader sample value = Authorization: Bearer <token>
// tokenString and token sample value = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30
func validateAndParse(authenticationHeader, key string) (any, error) {
	token, err := validate(authenticationHeader)
	if err != nil {
		return nil, err
	}

	return parseClaim(token, key), nil
}

// INPUT: Authenticate: Bearer <token>
// OUTPUT: token
func extractTokenString(authenticationHeader string) (string, error) {
	if strings.TrimSpace(authenticationHeader) == "" {
		return "", errors.New("authorization header is required")
	}

	parts := strings.SplitN(authenticationHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}

func validate(authenticationHeader string) (*jwt.Token, error) {
	tokenString, err := extractTokenString(authenticationHeader)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return getSecretKey(), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	exp, ok := parseClaim(token, "exp").(float64)
	if !ok {
		return nil, errors.New("invalid exp token claims")
	}

	if time.Now().After(time.Unix(int64(exp), 0)) {
		return nil, errors.New("token expired")
	}

	return token, nil
}

func parseClaim(token *jwt.Token, key string) any {
	return token.Claims.(jwt.MapClaims)[key]
}

func getSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}
