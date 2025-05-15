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
		save(ctx *gin.Context)

		getAll(ctx *gin.Context)
		getAllBy(ctx *gin.Context)

		updateSubject(ctx *gin.Context)
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

func (c ControllerImpl) save(ctx *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "save failed " + err.Error(),
		})
		return
	}

	postRequest := struct {
		Subject string `json:"subject" binding:"required"`
		Content string `json:"content" binding:"required"`
	}{}

	if err := ctx.ShouldBind(&postRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "save failed " + err.Error(),
		})
		return
	}

	id, err := c.service.save(currentUserId, postRequest.Subject, postRequest.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "save failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, id)
}

func (c ControllerImpl) getAll(ctx *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	isDeleted, err := strconv.ParseBool(ctx.DefaultQuery("isDeleted", "false"))
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

	pagedPosts, err := c.service.getAll(currentUserId, isDeleted, pageRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, pagedPosts)
}

func (c ControllerImpl) getAllBy(ctx *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "get all by failed " + err.Error(),
		})
		return
	}

	isDeleted, err := strconv.ParseBool(ctx.DefaultQuery("isDeleted", "false"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all by failed " + err.Error(),
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

	posts, err := c.service.getAllBy(currentUserId, isDeleted, pageRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all by failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (c ControllerImpl) updateSubject(ctx *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "update subject failed " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "update subject failed " + err.Error(),
		})
		return
	}

	subject := ctx.Query("subject")
	_, err = c.service.updateSubject(currentUserId, postId, subject)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "update subject failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, subject)
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
			"message": "update content failed" + err.Error(),
		})
		return
	}

	content := ctx.Query("content")
	_, err = c.service.updateContent(currentUserId, postId, content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "update content failed" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, content)
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

	attachment := ctx.Query("attachment")
	_, err = c.service.updateAttachment(currentUserId, postId, attachment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "update attachment failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, attachment)
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

	_, err = c.service.deleteById(currentUserId, postId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, postId)
}
