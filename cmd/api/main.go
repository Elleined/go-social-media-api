package main

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"social-media-application/internal/comment"
	cr "social-media-application/internal/comment/reaction"
	"social-media-application/internal/emoji"
	"social-media-application/internal/post"
	pr "social-media-application/internal/post/reaction"
	"social-media-application/internal/refresh"
	"social-media-application/internal/user"
	mw "social-media-application/middlewares"
	"social-media-application/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	db, err := utils.InitMySQLConnection()
	if err != nil {
		panic("can't connect to database")
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
		log.Fatal("can't set trusted proxies")
		return
	}

	// Initialize middlewares
	r.Use(mw.SecurityHeaders)

	// Initialize refresh token module
	refreshRepository := refresh.NewRepository(db)
	refreshService := refresh.NewService(refreshRepository)
	refreshController := refresh.NewController(refreshService)
	refreshController.RegisterRoutes(r)

	// Initialize user module
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService, refreshService)
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
	postReactionRepository := pr.NewRepository(db)
	postReactionService := pr.NewService(postReactionRepository)
	postReactionController := pr.NewController(postReactionService)
	postReactionController.RegisterRoutes(r)

	// Initialize comment module
	commentRepository := comment.NewRepository(db)
	commentService := comment.NewService(commentRepository)
	commentController := comment.NewController(commentService)
	commentController.RegisterRoutes(r)

	// Initialize comment reaction module
	commentReactionRepository := cr.NewRepository(db)
	commentReactionService := cr.NewService(commentReactionRepository)
	commentReactionController := cr.NewController(commentReactionService)
	commentReactionController.RegisterRoutes(r)

	err = r.Run(os.Getenv("PORT"))
	if err != nil {
		panic("cannot start server" + err.Error())
	}
}
