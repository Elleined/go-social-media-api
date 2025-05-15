package comment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-application/internal/paging"
	"social-media-application/utils"
	"strconv"
)

type (
	Controller interface {
		save(ctx *gin.Context)

		getAll(ctx *gin.Context)

		updateContent(ctx *gin.Context)
		updateAttachment(ctx *gin.Context)

		deleteById(ctx *gin.Context)

		RegisterRoutes(e *gin.Engine)
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

func (c ControllerImpl) RegisterRoutes(e *gin.Engine) {
	r := e.Group("/users/posts/:id/comments")
	{
		r.POST("", c.save)

		r.GET("", c.getAll)

		r.PATCH("/:commentId/content", c.updateContent)
		r.PATCH("/:commentId/attachment", c.updateAttachment)

		r.DELETE("/:commentId", c.deleteById)
	}
}

func (c ControllerImpl) save(ctx *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	content := ctx.Query("content")

	id, err := c.service.save(currentUserId, postId, content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, id)
}

func (c ControllerImpl) getAll(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	isDeleted, err := strconv.ParseBool(ctx.DefaultQuery("isDeleted", "false"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	pageRequest, err := paging.NewPageRequestStr(ctx.DefaultQuery("page", "1"), ctx.DefaultQuery("pageSize", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	comments, err := c.service.getAll(postId, isDeleted, pageRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (c ControllerImpl) updateContent(ctx *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "update content failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "update content failed " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "update content failed " + err.Error(),
		})
		return
	}

	newContent := ctx.Query("newContent")

	_, err = c.service.updateContent(currentUserId, postId, commentId, newContent)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "update content failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, newContent)
}

func (c ControllerImpl) updateAttachment(ctx *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "update attachment failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "update attachment failed " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "update attachment failed " + err.Error(),
		})
		return
	}

	newAttachment := ctx.Query("newAttachment")

	_, err = c.service.updateAttachment(currentUserId, postId, commentId, newAttachment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "update attachment failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, newAttachment)
}

func (c ControllerImpl) deleteById(ctx *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	_, err = c.service.deleteById(currentUserId, postId, commentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
