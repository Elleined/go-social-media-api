package commentreaction

import "github.com/gin-gonic/gin"

type (
	Controller interface {
		save(e *gin.Context)

		getAll(e *gin.Context)
		getAllByEmoji(e *gin.Context)

		update(e *gin.Context)

		delete(e *gin.Context)

		RegisterRoutes(r *gin.Engine)
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

}

func (c ControllerImpl) save(e *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c ControllerImpl) getAll(e *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c ControllerImpl) getAllByEmoji(e *gin.Context) {
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
