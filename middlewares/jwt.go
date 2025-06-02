package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// authenticationHeader = Authorization: Bearer <jwt>
// tokenString referred as the jwt but its not validated yet
// token referred as the jwt and its validated

func JWT(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid authorization header format",
		})
		return
	}

	// Get the actual JWT without the Bearer as prefix: eyJhb...
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if strings.TrimSpace(tokenString) == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid authorization header format",
		})
		return
	}

	// Validate signing method and return secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return getSecretKey(), nil
	})
	if err != nil || !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}

	// Get the JWT expiration
	exp, ok := token.Claims.(jwt.MapClaims)["exp"].(float64)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid exp token claims",
		})
		return
	}

	// Validate if token is not expired
	if time.Now().After(time.Unix(int64(exp), 0)) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "token expired",
		})
		return
	}

	// Parse all claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Set("sub", claims["sub"])
		ctx.Set("iat", claims["iat"])
		ctx.Set("exp", claims["exp"])
		ctx.Set("iss", claims["iss"])
		ctx.Set("aud", claims["aud"])
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid claims",
		})
		return
	}

	ctx.Next()
}

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

func GetSubject(ctx *gin.Context) (int, error) {
	sub, exists := ctx.Get("sub")
	if !exists {
		return 0, errors.New("sub not found")
	}

	id, ok := sub.(float64)
	if !ok {
		return 0, errors.New("id is not a number")
	}

	return int(id), nil
}

func getSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}
