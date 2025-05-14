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
		save(e *gin.Context)

		getAll(e *gin.Context)

		updateContent(e *gin.Context)
		updateAttachment(e *gin.Context)

		deleteById(e *gin.Context)

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

	content := e.Query("content")

	id, err := c.service.save(currentUserId, postId, content)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, id)
}

func (c ControllerImpl) getAll(e *gin.Context) {
	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	page := e.DefaultQuery("page", "1")
	pageSize := e.DefaultQuery("pageSize", "10")

	limit, offset, err := paging.Paginate(page, pageSize)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	isDeleted, err := strconv.ParseBool(e.DefaultQuery("isDeleted", "false"))
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	comments, err := c.service.getAll(postId, isDeleted, limit, offset)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, comments)
}

func (c ControllerImpl) updateContent(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "update content failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "update content failed " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(e.Param("commentId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "update content failed " + err.Error(),
		})
		return
	}

	newContent := e.Query("newContent")

	_, err = c.service.updateContent(currentUserId, postId, commentId, newContent)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "update content failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, newContent)
}

func (c ControllerImpl) updateAttachment(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "update attachment failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "update attachment failed " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(e.Param("commentId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "update attachment failed " + err.Error(),
		})
		return
	}

	newAttachment := e.Query("newAttachment")

	_, err = c.service.updateAttachment(currentUserId, postId, commentId, newAttachment)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "update attachment failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, newAttachment)
}

func (c ControllerImpl) deleteById(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(e.Param("commentId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	_, err = c.service.deleteById(currentUserId, postId, commentId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusNoContent, nil)
}
