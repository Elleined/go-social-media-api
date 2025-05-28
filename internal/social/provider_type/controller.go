package provider_type

import "github.com/gin-gonic/gin"

type (
	Controller interface {
		save(ctx *gin.Context)

		getById(ctx *gin.Context)
		getAll(ctx *gin.Context)

		update(ctx *gin.Context)
		delete(ctx *gin.Context)

		RegisterRoutes(c *gin.Engine)
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
	r := e.Group("/provider_types")
	{
		r.POST("", c.save)
		r.GET("/:id", c.getById)
		r.GET("", c.getAll)

		r.PATCH("/:id", c.update)

		r.DELETE("/:id", c.delete)
	}
}

func (c ControllerImpl) save(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c ControllerImpl) getById(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c ControllerImpl) getAll(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c ControllerImpl) update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c ControllerImpl) delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
