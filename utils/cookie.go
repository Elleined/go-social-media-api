package utils

import (
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

// path: Control which endpoint should the cookie only accessible
// For example with path as "" will work on all endpoints
// For example with path as "/api" will work on "/api" and "/api/*"

// domain: when "" is supplied only works on current domain and does not work on any other subdomain:
// For example with domain as "" it will only work on current domain example.com
// For example with domain as "example.com" it will work on domains api.example.com, app.example.com

// secure: only sent with https not http
// httpOnly: cannot be access via JS

const (
	secure   = true
	httpOnly = true
)

func SetRefreshToken(ctx *gin.Context, value string) error {
	expirationInDays, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION_IN_DAYS"))
	if err != nil {
		return err
	}

	expirationInHours := time.Duration(expirationInDays) * 24 * time.Hour
	ctx.SetCookie("refreshToken", value, int(expirationInHours.Seconds()), "/", "", secure, httpOnly)
	return nil
}

func ClearRefreshToken(ctx *gin.Context) {
	ctx.SetCookie(
		"refreshToken", // cookie name
		"",             // value
		-1,             // maxAge negative to delete
		"/",            // path
		"",             // domain (empty = current domain)
		secure,         // secure
		httpOnly,       // httpOnly
	)
}

func SetAccessToken(ctx *gin.Context, value string) error {
	expirationInMinute, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_IN_MINUTE"))
	if err != nil {
		return err
	}

	expirationInSeconds := time.Duration(expirationInMinute).Seconds()
	ctx.SetCookie("accessToken", value, int(expirationInSeconds), "/", "", secure, httpOnly)
	return nil
}

func ClearAccessToken(ctx *gin.Context) {
	ctx.SetCookie(
		"accessToken", // cookie name
		"",            // value
		-1,            // maxAge negative to delete
		"/",           // path
		"",            // domain (empty = current domain)
		secure,        // secure
		httpOnly,      // httpOnly
	)
}
