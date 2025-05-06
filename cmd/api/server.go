package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"social-media-application/internal/post"
	"social-media-application/internal/post/reaction"
	"social-media-application/internal/user"
	mw "social-media-application/middlewares"
	"social-media-application/utils"
)

func init() {
	// Initialize godotenv
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic().Msg("cannot load the .env file for mysql")
		return
	}
}

func main() {
	port := ":8000"

	// Initialize gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Initialize Database Connection
	db, err := utils.GetConnection()
	if err != nil {
		log.Panic().Msg("cannot connect to database")
	}

	// root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Only trust API calls from localhost
	err = r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot set trusted proxies")
		return
	}

	// Initialize middlewares
	r.Use(mw.SecurityHeaders)

	// Initialize user module
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService)
	userController.RegisterRoutes(r)

	// Initialize post module
	postRepository := post.NewRepository(db)
	postService := post.NewService(postRepository)
	postController := post.NewController(postService)
	postController.RegisterRoutes(r)

	// Initialize post reaction module
	postReactionRepository := reaction.NewRepository(db)
	postReactionService := reaction.NewService(postReactionRepository)
	postReactionController := reaction.NewController(postReactionService)
	postReactionController.RegisterRoutes(r)

	err = r.Run(port)
	if err != nil {
		panic("cannot start server" + err.Error())
	}
}
