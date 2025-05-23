package refresh

import "github.com/gin-gonic/gin"

type (
	Controller interface {
		refresh(ctx *gin.Context)
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

func (c *ControllerImpl) RegisterRoutes(e *gin.Engine) {
	r := e.Group("/users/refresh")
	{
		r.POST("", c.refresh)
	}
}

func (c *ControllerImpl) refresh(ctx *gin.Context) {

}
