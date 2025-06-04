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

func SetRefreshTokenCookie(ctx *gin.Context, value string) {
	refreshTokenExpiryDays, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION_IN_DAYS"))
	if err != nil {
		return
	}

	refreshTokenExpiryHour := time.Duration(refreshTokenExpiryDays) * 24 * time.Hour
	ctx.SetCookie("refreshToken", value, int(refreshTokenExpiryHour.Seconds()), "/", "", false, true)
}
