package refresh

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-application/middlewares"
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
	e.POST("/users/refresh-tokens", c.refresh)

	r := e.Group("/users/refresh-tokens", middleware.JWT)
	{
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
	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "refresh failed " + err.Error(),
		})
		return
	}

	// get the refresh token from database
	oldRefreshToken, err := c.service.getBy(refreshToken)
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
	accessToken, err := middleware.GenerateJWT(oldRefreshToken.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "refresh failed! " + err.Error(),
		})
		return
	}

	err = utils.SetRefreshTokenCookie(ctx, newRefreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "refresh failed! " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, accessToken)
}

func (c *ControllerImpl) getAllBy(ctx *gin.Context) {
	sub, err := middleware.GetSubject(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get all by failed " + err.Error(),
		})
		return
	}

	refreshTokens, err := c.service.getAllBy(sub)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all by failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, refreshTokens)
}

func (c *ControllerImpl) revoke(ctx *gin.Context) {
	sub, err := middleware.GetSubject(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "save failed " + err.Error(),
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
	_, err = c.service.revoke(id, sub)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "revoke failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
