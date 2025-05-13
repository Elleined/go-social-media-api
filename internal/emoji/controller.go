package emoji

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type (
	Controller interface {
		save(e *gin.Context)
		getAll(e *gin.Context)
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
	r := e.Group("/emojis")
	{
		r.POST("", c.save)
		r.GET("", c.getAll)
		r.PATCH("/:id", c.update)
		r.DELETE("/:id", c.delete)
	}
}

func (c ControllerImpl) save(e *gin.Context) {
	name := e.Query("name")

	id, err := c.service.save(name)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "saved failed" + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, id)
}

func (c ControllerImpl) getAll(e *gin.Context) {
	emojis, err := c.service.getAll()
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, emojis)
}

func (c ControllerImpl) update(e *gin.Context) {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "update failed " + err.Error(),
		})
		return
	}

	name := e.Query("name")
	_, err = c.service.update(id, name)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "update failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, name)
}

func (c ControllerImpl) delete(e *gin.Context) {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	_, err = c.service.delete(id)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusNoContent, nil)
}
