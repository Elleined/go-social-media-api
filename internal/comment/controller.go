package comment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-application/utils"
	"strconv"
)

type Controller interface {
	save(e *gin.Context)

	getAll(e *gin.Context)

	updateContent(e *gin.Context)
	updateAttachment(e *gin.Context)

	deleteById(e *gin.Context)

	RegisterRoutes(e *gin.Engine)
}

type ControllerImpl struct {
	service Service
}

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
		r.PATCH("/:commentId/attachment", c.getAll)

		r.DELETE("/:commentId", c.deleteById)
	}
}

func (c ControllerImpl) save(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "something wrong with jwt " + err.Error(),
		})
		return
	}

	commentRequest := struct {
		PostId  int    `json:"post_id" binding:"required"`
		Content string `json:"content"  binding:"required"`
	}{}

	if err := e.ShouldBindJSON(&commentRequest); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't save comment malformed request" + err.Error(),
		})
		return
	}

	id, err := c.service.save(currentUserId, commentRequest.PostId, commentRequest.Content)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't save comment " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, id)
}

func (c ControllerImpl) getAll(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "something wrong with jwt " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't get all comment " + err.Error(),
		})
		return
	}

	page := e.DefaultQuery("page", "1")
	pageSize := e.DefaultQuery("pageSize", "10")

	limit, offset, err := utils.Paginate(page, pageSize)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't get all comment " + err.Error(),
		})
		return
	}

	comments, err := c.service.getAll(currentUserId, postId, limit, offset)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't get all comment " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, comments)
}

func (c ControllerImpl) updateContent(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "something wrong with jwt " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't update comment content " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(e.Param("commentId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't update comment content " + err.Error(),
		})
		return
	}

	newContent := e.Query("newContent")

	_, err = c.service.updateContent(currentUserId, postId, commentId, newContent)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't update comment content " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, newContent)
}

func (c ControllerImpl) updateAttachment(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "something wrong with jwt " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't update comment attachment " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(e.Param("commentId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't update comment attachment " + err.Error(),
		})
		return
	}

	newAttachment := e.Query("newAttachment")

	_, err = c.service.updateAttachment(currentUserId, postId, commentId, newAttachment)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't update comment attachment " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, newAttachment)
}

func (c ControllerImpl) deleteById(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "something wrong with jwt " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't delete comment " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(e.Param("commentId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't delete comment " + err.Error(),
		})
		return
	}

	_, err = c.service.deleteById(currentUserId, postId, commentId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't delete comment " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusNoContent, commentId)
}
