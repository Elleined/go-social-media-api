package emoji

import (
	"github.com/gin-gonic/gin"
	"net/http"
	middleware "social-media-application/middlewares"
	"strconv"
)

type (
	Controller interface {
		save(ctx *gin.Context)

		getById(ctx *gin.Context)
		getAll(ctx *gin.Context)
		update(ctx *gin.Context)
		delete(ctx *gin.Context)

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
	r := e.Group("/emojis", middleware.JWT)
	{
		r.POST("", c.save)
		r.GET("/:id", c.getById)
		r.GET("", c.getAll)
		r.PATCH("/:id", c.update)
		r.DELETE("/:id", c.delete)
	}
}

func (c ControllerImpl) save(ctx *gin.Context) {
	name := ctx.Query("name")

	id, err := c.service.save(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "saved failed" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, id)
}

func (c ControllerImpl) getById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get by id failed " + err.Error(),
		})
		return
	}

	emoji, err := c.service.getById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get by id failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, emoji)
}

func (c ControllerImpl) getAll(ctx *gin.Context) {
	emojis, err := c.service.getAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, emojis)
}

func (c ControllerImpl) update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "update failed " + err.Error(),
		})
		return
	}

	name := ctx.Query("name")
	_, err = c.service.update(id, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "update failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, name)
}

func (c ControllerImpl) delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	_, err = c.service.delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
