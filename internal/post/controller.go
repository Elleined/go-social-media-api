package post

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
		getAllBy(e *gin.Context)

		updateSubject(e *gin.Context)
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
	r := e.Group("/users/posts")
	{
		r.POST("", c.save)

		r.GET("", c.getAll)
		r.GET("/all-by-user", c.getAllBy)

		r.PATCH("/:id/subject", c.updateSubject)
		r.PATCH("/:id/content", c.updateContent)
		r.PATCH("/:id/attachment", c.updateAttachment)

		r.DELETE("/:id", c.deleteById)
	}
}

func (c ControllerImpl) save(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "save failed " + err.Error(),
		})
		return
	}

	postRequest := struct {
		Subject string `json:"subject" binding:"required"`
		Content string `json:"content" binding:"required"`
	}{}

	if err := e.BindJSON(&postRequest); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "save failed " + err.Error(),
		})
		return
	}

	id, err := c.service.save(currentUserId, postRequest.Subject, postRequest.Content)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "save failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, id)
}

func (c ControllerImpl) getAll(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	page, err := strconv.Atoi(e.DefaultQuery("page", "1"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	pageSize, err := strconv.Atoi(e.DefaultQuery("pageSize", "10"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	isDeleted, err := strconv.ParseBool(e.DefaultQuery("isDeleted", "false"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	pageRequest, err := paging.NewPageRequest(page, pageSize)
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	pagedPosts, err := c.service.getAll(currentUserId, isDeleted, pageRequest)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, pagedPosts)
}

func (c ControllerImpl) getAllBy(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "get all by failed " + err.Error(),
		})
		return
	}

	page, err := strconv.Atoi(e.DefaultQuery("page", "1"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	pageSize, err := strconv.Atoi(e.DefaultQuery("pageSize", "10"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	isDeleted, err := strconv.ParseBool(e.DefaultQuery("isDeleted", "false"))
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all by failed " + err.Error(),
		})
		return
	}

	pageRequest, err := paging.NewPageRequest(page, pageSize)
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	posts, err := c.service.getAllBy(currentUserId, isDeleted, pageRequest)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all by failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, posts)
}

func (c ControllerImpl) updateSubject(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "update subject failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "update subject failed " + err.Error(),
		})
		return
	}

	subject := e.Query("subject")
	_, err = c.service.updateSubject(currentUserId, postId, subject)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "update subject failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, subject)
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
			"message": "update content failed" + err.Error(),
		})
		return
	}

	content := e.Query("content")
	_, err = c.service.updateContent(currentUserId, postId, content)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "update content failed" + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, content)
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

	attachment := e.Query("attachment")
	_, err = c.service.updateAttachment(currentUserId, postId, attachment)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "update attachment failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, attachment)
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

	_, err = c.service.deleteById(currentUserId, postId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusNoContent, postId)
}
