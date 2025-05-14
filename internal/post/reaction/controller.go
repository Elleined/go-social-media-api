package reaction

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-application/internal/paging"
	"social-media-application/utils"
	"strconv"
)

type (
	Controller interface {
		save(e *gin.Context)

		findAll(e *gin.Context)
		findAllByEmoji(e *gin.Context)

		update(e *gin.Context)

		delete(e *gin.Context)

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
	r := e.Group("/users/posts/:id/reactions")
	{
		r.POST("", c.save)

		r.GET("", c.findAll)
		r.GET("/emoji/:emojiId", c.findAllByEmoji)

		r.PATCH("/emoji/:emojiId", c.update)
		r.DELETE("", c.delete)
	}
}

func (c ControllerImpl) save(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	emojiId, err := strconv.Atoi(e.Query("emojiId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	id, err := c.service.save(currentUserId, postId, emojiId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, id)
}

func (c ControllerImpl) findAll(e *gin.Context) {
	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	pageRequest, err := paging.NewPageRequestStr(e.DefaultQuery("page", "1"), e.DefaultQuery("pageSize", "10"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	reactions, err := c.service.getAll(postId, pageRequest)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, reactions)
}

func (c ControllerImpl) findAllByEmoji(e *gin.Context) {
	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all by emoji failed " + err.Error(),
		})
		return
	}

	emojiId, err := strconv.Atoi(e.Param("emojiId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all by emoji failed " + err.Error(),
		})
		return
	}

	pageRequest, err := paging.NewPageRequestStr(e.DefaultQuery("page", "1"), e.DefaultQuery("pageSize", "10"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	reactions, err := c.service.getAllByEmoji(postId, emojiId, pageRequest)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all by emoji failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, reactions)
}

func (c ControllerImpl) update(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "update failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "update failed " + err.Error(),
		})
		return
	}

	newEmojiId, err := strconv.Atoi(e.Param("emojiId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "update failed " + err.Error(),
		})
		return
	}

	_, err = c.service.update(currentUserId, postId, newEmojiId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "update failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, newEmojiId)
}

func (c ControllerImpl) delete(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "delete failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "delete failed " + err.Error(),
		})
		return
	}

	_, err = c.service.delete(currentUserId, postId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusNoContent, nil)
}
