package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-application/utils"
	"strconv"
)

type Controller interface {
	save(e *gin.Context)

	getById(e *gin.Context)
	getByEmail(e *gin.Context)

	getAll(e *gin.Context)

	deleteById(e *gin.Context)

	changeStatus(e *gin.Context)
	changePassword(e *gin.Context)

	login(e *gin.Context)

	RegisterRoutes(c *gin.Engine)
}

type ControllerImpl struct {
	service Service
}

func NewController(service Service) Controller {
	return &ControllerImpl{
		service: service,
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

func (c *ControllerImpl) save(e *gin.Context) {
	userRequest := struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
	}{}

	if err := e.ShouldBind(&userRequest); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't save user! malformed request" + err.Error(),
		})
		return
	}

	id, err := c.service.save(userRequest.FirstName, userRequest.LastName, userRequest.Email, userRequest.Password)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't save user " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusCreated, id)
}

func (c *ControllerImpl) getById(e *gin.Context) {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id " + err.Error(),
		})
		return
	}

	user, err := c.service.getById(id)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't get by id " + err.Error(),
		})
	}

	e.JSON(http.StatusOK, user)
}

func (c *ControllerImpl) getByEmail(e *gin.Context) {
	email := e.Param("email")

	user, err := c.service.getByEmail(email)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't get by email " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, user)
}

func (c *ControllerImpl) getAll(e *gin.Context) {
	isActive := e.DefaultQuery("isActive", "true")
	page := e.DefaultQuery("page", "1")
	pageSize := e.DefaultQuery("pageSize", "10")

	limit, offset, err := utils.Paginate(page, pageSize)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "something wrong with pagination " + err.Error(),
		})
		return
	}

	isActiveBool, err := strconv.ParseBool(isActive)
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid isActive " + err.Error(),
		})
		return
	}

	users, err := c.service.getAll(isActiveBool, limit, offset)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot fetch all users " + err.Error(),
		})
	}

	e.JSON(http.StatusOK, users)
}

func (c *ControllerImpl) deleteById(e *gin.Context) {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id " + err.Error(),
		})
		return
	}

	_, err = c.service.deleteById(id)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't delete " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, nil)
}

func (c *ControllerImpl) changeStatus(e *gin.Context) {
	status, err := strconv.ParseBool(e.Query("status"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid status " + err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id " + err.Error(),
		})
		return
	}

	_, err = c.service.changeStatus(id, status)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't change status " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, status)
}

func (c *ControllerImpl) changePassword(e *gin.Context) {
	passwordRequest := struct {
		Password string
	}{}
	if err := e.ShouldBind(&passwordRequest); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid password " + err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id " + err.Error(),
		})
		return
	}

	_, err = c.service.changePassword(id, passwordRequest.Password)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't change password " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, id)
}

func (c *ControllerImpl) login(e *gin.Context) {
	loginRequest := struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := e.ShouldBind(&loginRequest); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid malformed request " + err.Error(),
		})
		return
	}

	jwt, err := c.service.login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't login " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, jwt)
}
