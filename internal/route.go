package internal

import "github.com/gin-gonic/gin"

type Route interface {
	RegisterRoutes(e *gin.Engine)
}
