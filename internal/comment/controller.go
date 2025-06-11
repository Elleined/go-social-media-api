package comment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-application/internal/paging"
	"social-media-application/middlewares"
	"strconv"
)

type (
	Controller interface {
		save(ctx *gin.Context)

		getById(ctx *gin.Context)
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
	r := e.Group("/users/posts/:id/comments", middleware.JWT)
	{
		r.POST("", c.save)

		r.GET("/:commentId", c.getById)
		r.GET("", c.getAll)

		r.PATCH("/:commentId/content", c.updateContent)
		r.PATCH("/:commentId/attachment", c.updateAttachment)

		r.DELETE("/:commentId", c.deleteById)
	}
}

func (c ControllerImpl) save(ctx *gin.Context) {
	sub, err := middleware.GetSubject(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "save failed " + err.Error(),
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

	request := struct {
		Content    string `json:"content" binding:"required"`
		Attachment string `json:"attachment"`
	}{}
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "save failed " + err.Error(),
		})
		return
	}

	id, err := c.service.save(sub, postId, request.Content, request.Attachment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, id)
}

func (c ControllerImpl) getById(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get by id failed " + err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get by id failed " + err.Error(),
		})
		return
	}

	comment, err := c.service.getById(postId, commentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get by id failed " + err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, comment)
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

	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	field := ctx.DefaultQuery("field", "created_at")
	sortBy := ctx.DefaultQuery("sortBy", "DESC")
	request, err := paging.NewPageRequestStr(page, pageSize, field, sortBy)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	comments, err := c.service.getAll(postId, isDeleted, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (c ControllerImpl) updateContent(ctx *gin.Context) {
	sub, err := middleware.GetSubject(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "save failed " + err.Error(),
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

	content := ctx.Query("content")
	_, err = c.service.updateContent(sub, postId, commentId, content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "update content failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, content)
}

func (c ControllerImpl) updateAttachment(ctx *gin.Context) {
	sub, err := middleware.GetSubject(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "save failed " + err.Error(),
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

	attachment := ctx.Query("attachment")
	_, err = c.service.updateAttachment(sub, postId, commentId, attachment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "update attachment failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, attachment)
}

func (c ControllerImpl) deleteById(ctx *gin.Context) {
	sub, err := middleware.GetSubject(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "save failed " + err.Error(),
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

	_, err = c.service.deleteById(sub, postId, commentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
