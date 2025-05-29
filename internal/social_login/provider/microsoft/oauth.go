package microsoft

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
	"io"
	"net/http"
	"os"
	"social-media-application/internal/refresh"
	"social-media-application/internal/social_login"
	"social-media-application/internal/user"
)

func InitMSLogin() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  os.Getenv("MICROSOFT_REDIRECT_URL"),
		ClientID:     os.Getenv("MICROSOFT_KEY"),
		ClientSecret: os.Getenv("MICROSOFT_SECRET"),
		Scopes:       []string{"User.Read"},
		Endpoint:     microsoft.AzureADEndpoint(os.Getenv("MICROSOFT_TENANT_ID")),
	}
}

type Controller struct {
	config            *oauth2.Config
	refreshService    refresh.Service
	socialUserService social_login.Service
	userService       user.Service
}

func NewController(config *oauth2.Config, refreshService refresh.Service, socialUserService social_login.Service, userService user.Service) *Controller {
	return &Controller{
		config:            config,
		refreshService:    refreshService,
		socialUserService: socialUserService,
		userService:       userService,
	}
}

func (c Controller) RegisterRoutes(e *gin.Engine) {
	r := e.Group("/auth/microsoft")
	{
		r.GET("", c.login)
		r.GET("/callback", c.callback)
	}
}

func (c Controller) login(ctx *gin.Context) {
	// Redirect user to Microsoft login page
	url := c.config.AuthCodeURL(
		"state",
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam("prompt", "login"),
	)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (c Controller) callback(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing code",
		})
		return
	}

	token, err := c.config.Exchange(ctx.Request.Context(), code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "authentication failed " + err.Error(),
		})
		return
	}

	client := c.config.Client(ctx.Request.Context(), token)
	resp, err := client.Get("https://graph.microsoft.com/v1.0/me")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get user information " + err.Error(),
		})
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	var userInfo any
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to parse user info",
		})
		return
	}

	ctx.JSON(http.StatusOK, userInfo)
}
