package commentreaction

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
		getAllByEmoji(ctx *gin.Context)

		update(ctx *gin.Context)

		delete(ctx *gin.Context)

		RegisterRoutes(r *gin.Engine)
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
	r := e.Group("/users/posts/:id/comments/:commentId/reactions")
	{
		r.POST("", c.save)

		r.GET("", c.getAll)
		r.GET("/emoji/:emojiId", c.getAllByEmoji)

		r.PATCH("/emoji/:emojiId`", c.update)
		r.DELETE("", c.delete)
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

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	emojiId, err := strconv.Atoi(ctx.Query("emojiId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	id, err := c.service.save(currentUserId, postId, commentId, emojiId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (c ControllerImpl) getAll(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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

	reactions, err := c.service.findAll(postId, commentId, pageRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, reactions)
}

func (c ControllerImpl) getAllByEmoji(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get all by emoji failed " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get all by emoji failed " + err.Error(),
		})
		return
	}

	emojiId, err := strconv.Atoi(ctx.Param("emojiId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get all by emoji failed " + err.Error(),
		})
		return
	}

	pageRequest, err := paging.NewPageRequestStr(ctx.DefaultQuery("page", "1"), ctx.DefaultQuery("pageSize", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get all by emoji failed " + err.Error(),
		})
		return
	}

	reactions, err := c.service.findAllByEmoji(postId, commentId, emojiId, pageRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all by emoji failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, reactions)
}

func (c ControllerImpl) update(ctx *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "get all by emoji failed " + err.Error(),
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

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	emojiId, err := strconv.Atoi(ctx.Param("emojiId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	_, err = c.service.update(currentUserId, postId, commentId, emojiId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, emojiId)
}

func (c ControllerImpl) delete(ctx *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "delete failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "delete failed " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "delete failed " + err.Error(),
		})
		return
	}

	_, err = c.service.delete(currentUserId, postId, commentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
