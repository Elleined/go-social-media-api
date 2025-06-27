package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func SetTokens(ctx *gin.Context, accessToken string, refreshToken string) error {
	err := setRefreshToken(ctx, refreshToken)
	if err != nil {
		return err
	}

	err = setAccessToken(ctx, accessToken)
	if err != nil {
		return err
	}

	return nil
}

func ClearTokens(ctx *gin.Context) {
	clearAccessToken(ctx)
	clearRefreshToken(ctx)
}

func setRefreshToken(ctx *gin.Context, value string) error {
	expirationInDays, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION_IN_DAYS"))
	if err != nil {
		return err
	}

	expirationInHours := time.Duration(expirationInDays) * 24 * time.Hour
	ctx.SetCookie("refreshToken", value, int(expirationInHours.Seconds()), "/", "", secure, httpOnly)
	ctx.SetSameSite(http.SameSiteStrictMode)
	return nil
}

func clearRefreshToken(ctx *gin.Context) {
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

func setAccessToken(ctx *gin.Context, value string) error {
	expirationInMinute, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_IN_MINUTE"))
	if err != nil {
		return err
	}

	expirationInSeconds := time.Duration(expirationInMinute).Seconds()
	ctx.SetCookie("accessToken", value, int(expirationInSeconds), "/", "", secure, httpOnly)
	ctx.SetSameSite(http.SameSiteStrictMode)
	return nil
}

func clearAccessToken(ctx *gin.Context) {
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
