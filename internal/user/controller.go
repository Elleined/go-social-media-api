package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-application/internal/paging"
	"social-media-application/internal/refresh"
	pd "social-media-application/internal/user/password"
	"social-media-application/middlewares"
	"strconv"
)

type (
	Controller interface {
		save(ctx *gin.Context)

		getById(ctx *gin.Context)
		getByEmail(ctx *gin.Context)

		getAll(ctx *gin.Context)

		deleteById(ctx *gin.Context)

		changeStatus(ctx *gin.Context)
		changePassword(ctx *gin.Context)

		login(ctx *gin.Context)

		RegisterRoutes(c *gin.Engine)
	}

	ControllerImpl struct {
		service        Service
		refreshService refresh.Service
	}
)

func NewController(service Service, refreshService refresh.Service) Controller {
	return &ControllerImpl{
		service:        service,
		refreshService: refreshService,
	}
}

func (c *ControllerImpl) RegisterRoutes(e *gin.Engine) {
	r := e.Group("/users")
	{
		r.POST("/login", c.login)
		r.POST("", c.save)

		r.GET("/id/:id", c.getById)
		r.GET("/email/:email", c.getByEmail)
		r.GET("", c.getAll)

		r.DELETE("/:id", c.deleteById)

		r.PATCH("/:id/status", c.changeStatus)
		r.PATCH("/:id/password", c.changePassword)
	}
}

func (c *ControllerImpl) save(ctx *gin.Context) {
	userRequest := struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
	}{}

	if err := ctx.ShouldBind(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	id, err := c.service.save(userRequest.FirstName, userRequest.LastName, userRequest.Email, userRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "saved failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (c *ControllerImpl) getById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "get by id failed " + err.Error(),
		})
		return
	}

	user, err := c.service.getById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get by id failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *ControllerImpl) getByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := c.service.getByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get by email failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *ControllerImpl) getAll(ctx *gin.Context) {
	isActive, err := strconv.ParseBool(ctx.DefaultQuery("isActive", "true"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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

	users, err := c.service.getAll(isActive, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get all failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *ControllerImpl) deleteById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	_, err = c.service.deleteById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete by id failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *ControllerImpl) changeStatus(ctx *gin.Context) {
	status, err := strconv.ParseBool(ctx.Query("status"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "change status failed " + err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "change status failed " + err.Error(),
		})
		return
	}

	_, err = c.service.changeStatus(id, status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "change status failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, status)
}

func (c *ControllerImpl) changePassword(ctx *gin.Context) {
	passwordRequest := struct {
		Password string
	}{}

	if err := ctx.ShouldBind(&passwordRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "change password failed " + err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "change password failed " + err.Error(),
		})
		return
	}

	_, err = c.service.changePassword(id, passwordRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "change password failed " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, id)
}

func (c *ControllerImpl) login(ctx *gin.Context) {
	loginRequest := struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := ctx.ShouldBind(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "login failed " + err.Error(),
		})
		return
	}

	user, err := c.service.getByEmail(loginRequest.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "login failed! invalid credentials ",
		})
		return
	}

	if !pd.IsPasswordMatch(loginRequest.Password, user.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "login failed! invalid credentials",
		})
		return
	}

	jwt, err := middleware.GenerateJWT(user.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "login failed! " + err.Error(),
		})
		return
	}

	token, err := c.refreshService.Save(user.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "login failed! " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"refresh_token": token,
		"access_token":  jwt,
		"message":       "saved the refresh token securely",
	})
}
