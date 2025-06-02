package refresh

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-application/utils"
	"strconv"
)

type (
	Controller interface {
		refresh(ctx *gin.Context)
		getAllBy(ctx *gin.Context)
		revoke(ctx *gin.Context)
		RegisterRoutes(c *gin.Engine)
	}

	ControllerImpl struct {
		service Service
	}
)

func NewController(service Service) Controller {
	return &ControllerImpl{
		service: service,
	}
}

func (c *ControllerImpl) RegisterRoutes(e *gin.Engine) {
	r := e.Group("/users/refresh-tokens")
	{
		r.POST("", c.refresh)
		r.GET("", c.getAllBy)
		r.DELETE("/:id", c.revoke)
	}
}

// 1.isValid
// 2. revoke the old token
// 3. Generate and save the new refresh token and return it
// 4. Generate new access token and return it
func (c *ControllerImpl) refresh(ctx *gin.Context) {
	// get the refresh token from client
	request := struct {
		Token string `json:"refresh_token" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "refresh failed " + err.Error(),
		})
		return
	}

	// get the refresh token from database
	oldRefreshToken, err := c.service.getBy(request.Token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "refresh failed " + err.Error(),
		})
		return
	}

	// 1. isValid
	// checking if refresh token is not expired and not revoked
	err = c.service.isValid(oldRefreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "refresh failed! " + err.Error(),
		})
		return
	}

	// 2. revoke the old token
	_, err = c.service.revoke(oldRefreshToken.Id, oldRefreshToken.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "refresh failed! " + err.Error(),
		})
		return
	}

	// 3. Generate and save the new refresh token and return it
	newRefreshToken, err := c.service.SaveWith(oldRefreshToken.UserId, oldRefreshToken.ExpiresAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "refresh failed! " + err.Error(),
		})
		return
	}

	// 4. Generate new access token and return it
	accessToken, err := utils.GenerateJWT(oldRefreshToken.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "refresh failed! " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"refresh_token": newRefreshToken,
		"access_token":  accessToken,
		"message":       "saved the refresh token securely",
	})
}

func (c *ControllerImpl) getAllBy(ctx *gin.Context) {
	currentUserId, err := utils.GetSubject(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "get all by failed " + err.Error(),
		})
		return
	}

	refreshTokens, err := c.service.getAllBy(currentUserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all by failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, refreshTokens)
}

func (c *ControllerImpl) revoke(ctx *gin.Context) {
	currentUserId, err := utils.GetSubject(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "get all by failed " + err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "revoke failed " + err.Error(),
		})
		return
	}
	_, err = c.service.revoke(id, currentUserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "revoke failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
