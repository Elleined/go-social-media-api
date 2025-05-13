package emoji

import "github.com/gin-gonic/gin"

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
	//TODO implement me
	panic("implement me")
}

func (c ControllerImpl) getAll(e *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c ControllerImpl) update(e *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c ControllerImpl) delete(e *gin.Context) {
	//TODO implement me
	panic("implement me")
}
