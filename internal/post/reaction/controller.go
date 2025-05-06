package reaction

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-application/utils"
	"strconv"
)

type Controller interface {
	save(e *gin.Context)

	findAll(e *gin.Context)
	findAllByEmoji(e *gin.Context)

	delete(e *gin.Context)

	RegisterRoutes(e *gin.Engine)
}

type ControllerImpl struct {
	service Service
}

func NewController(service Service) Controller {
	return &ControllerImpl{
		service: service,
	}
}

func (c ControllerImpl) RegisterRoutes(e *gin.Engine) {
	r := e.Group("/users/posts/:id/reactions")
	{
		r.POST("", c.save)

		r.GET("", c.findAll)
		r.GET("/:emojiId", c.findAllByEmoji)

		r.DELETE("", c.delete)
	}
}

func (c ControllerImpl) save(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "something wrong with jwt " + err.Error(),
		})
		return
	}

	reactionRequest := struct {
		PostId  int `json:"post_id"`
		EmojiId int `json:"emoji_id"`
	}{}

	if err := e.BindJSON(&reactionRequest); err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't save reaction " + err.Error(),
		})
		return
	}

	id, err := c.service.save(currentUserId, reactionRequest.PostId, reactionRequest.EmojiId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't save reaction " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, id)
}

func (c ControllerImpl) findAll(e *gin.Context) {
	postId, err := strconv.Atoi(e.Param("postId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't find all posts " + err.Error(),
		})
		return
	}

	reactions, err := c.service.findAll(postId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't find all posts " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, reactions)
}

func (c ControllerImpl) findAllByEmoji(e *gin.Context) {
	postId, err := strconv.Atoi(e.Param("postId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't find all posts by emoji " + err.Error(),
		})
		return
	}

	emojiId, err := strconv.Atoi(e.Param("emojiId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't find all posts by emoji " + err.Error(),
		})
		return
	}

	reactions, err := c.service.findAllByEmoji(postId, emojiId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't find all posts by emoji " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, reactions)
}

func (c ControllerImpl) delete(e *gin.Context) {
	currentUserId, err := utils.GetCurrentUserId(e.GetHeader("Authorization"))
	if err != nil {
		e.JSON(http.StatusUnauthorized, gin.H{
			"message": "something wrong with jwt " + err.Error(),
		})
		return
	}

	postId, err := strconv.Atoi(e.Param("postId"))
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": "can't find all posts " + err.Error(),
		})
		return
	}

	_, err = c.service.delete(currentUserId, postId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't find all posts " + err.Error(),
		})
		return
	}

	e.JSON(http.StatusOK, gin.H{})
}
