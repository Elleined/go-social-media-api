package main

import (
	"github.com/jmoiron/sqlx"
	"os"
	"social-media-application/internal/comment"
	"social-media-application/internal/emoji"
	"social-media-application/internal/post"
	"social-media-application/internal/post/reaction"
	"social-media-application/internal/user"
	mw "social-media-application/middlewares"
	"social-media-application/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init() {
	ginMode := os.Getenv("GIN_MODE")

	// Only load the godotenv when running in debug mode
	// But in release mode the .env will be supplied dynamically
	if ginMode == gin.ReleaseMode || strings.TrimSpace(ginMode) == "" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}
}

func main() {
	// Initialize Database Connection
	db, err := utils.GetConnection()
	if err != nil {
		log.Panic().Msg("cannot connect to database")
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	// Initialize gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

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

	// Initialize emoji module
	emojiRepository := emoji.NewRepository(db)
	emojiService := emoji.NewService(emojiRepository)
	emojiController := emoji.NewController(emojiService)
	emojiController.RegisterRoutes(r)

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

	// Initialize comment module
	commentRepository := comment.NewRepository(db)
	commentService := comment.NewService(commentRepository)
	commentController := comment.NewController(commentService)
	commentController.RegisterRoutes(r)

	err = r.Run(os.Getenv("PORT"))
	if err != nil {
		panic("cannot start server" + err.Error())
	}
}
