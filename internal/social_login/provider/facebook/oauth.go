package facebook

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"io"
	"net/http"
	"os"
	"social-media-application/internal/refresh"
	"social-media-application/internal/social_login"
	"social-media-application/internal/social_login/provider_type"
	"social-media-application/internal/user"
	middleware "social-media-application/middlewares"
	"social-media-application/utils"
	"strings"
)

func InitFacebookLogin() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("FACEBOOK_KEY"),
		ClientSecret: os.Getenv("FACEBOOK_SECRET"),
		RedirectURL:  os.Getenv("FACEBOOK_REDIRECT_URL"),
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	}
}

type Controller struct {
	config              *oauth2.Config
	refreshService      refresh.Service
	socialUserService   social_login.Service
	userService         user.Service
	providerTypeService provider_type.Service
}

func NewController(config *oauth2.Config, refreshService refresh.Service, socialUserService social_login.Service, userService user.Service, providerTypeService provider_type.Service) *Controller {
	return &Controller{
		config:              config,
		refreshService:      refreshService,
		socialUserService:   socialUserService,
		userService:         userService,
		providerTypeService: providerTypeService,
	}
}

func (c Controller) RegisterRoutes(e *gin.Engine) {
	r := e.Group("/auth/facebook")
	{
		r.GET("", c.login)
		r.GET("/callback", c.callback)
	}
}

func (c Controller) login(ctx *gin.Context) {
	// Redirect user to Google login page
	ctx.Redirect(http.StatusTemporaryRedirect,
		c.config.AuthCodeURL(
			"state",
			oauth2.AccessTypeOffline,
			oauth2.SetAuthURLParam("prompt", "login"),
		))
}

func (c Controller) callback(ctx *gin.Context) {
	// Extract the "code" query parameter from the URL.
	code := ctx.Query("code")
	if strings.TrimSpace(code) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing code",
		})
		return
	}

	// Exchange the authorization code for an OAuth2 token.
	token, err := c.config.Exchange(ctx.Request.Context(), code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "authentication failed " + err.Error(),
		})
		return
	}

	// Create an HTTP client using the token.
	client := c.config.Client(ctx.Request.Context(), token)

	// Make a GET request to the social provider's user info endpoint.
	resp, err := client.Get("https://graph.facebook.com/me?fields=id,name,email&access_token=" + token.AccessToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get user information " + err.Error(),
		})
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	userInfo := struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse user info",
		})
		return
	}
	// end of auth provider call

	// Start of backend logic
	// Get provider type if exists
	providerType, err := c.providerTypeService.GetByName("FACEBOOK")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get provider type",
		})
		return
	}

	// 1 User already exists
	socialUser, err := c.socialUserService.GetByProviderTypeAndId(providerType.Id, userInfo.Id)
	if err == nil {
		accessToken, refreshToken, err := c.generateTokens(socialUser.UserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "login failed! " + err.Error(),
			})
			return
		}

		err = utils.SetTokens(ctx, accessToken, refreshToken)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "login failed! " + err.Error(),
			})
			return
		}

		ctx.Redirect(http.StatusFound, os.Getenv("FRONT_END_REDIRECT_URL"))
		return
	}

	// 2 User already exist and linked to other social account
	existingUser, err := c.userService.GetByEmail(userInfo.Email)
	if err == nil {
		// Means local user
		if strings.TrimSpace(existingUser.Password) != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Login failed! Please login locally",
			})
			return
		}

		_, err := c.socialUserService.Save(providerType.Id, existingUser.Id, userInfo.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "login failed! " + err.Error(),
			})
			return
		}

		accessToken, refreshToken, err := c.generateTokens(existingUser.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "login failed! " + err.Error(),
			})
			return
		}

		err = utils.SetTokens(ctx, accessToken, refreshToken)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "login failed! " + err.Error(),
			})
			return
		}

		ctx.Redirect(http.StatusFound, os.Getenv("FRONT_END_REDIRECT_URL"))
		return
	}

	// 3 User not exists and no links to other social account
	id, err := c.userService.SaveSocial(userInfo.Name, userInfo.Name, userInfo.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "login failed! " + err.Error(),
		})
		return
	}

	_, err = c.socialUserService.Save(providerType.Id, int(id), userInfo.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "login failed! " + err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := c.generateTokens(int(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "login failed! " + err.Error(),
		})
		return
	}

	err = utils.SetTokens(ctx, accessToken, refreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "login failed! " + err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusFound, os.Getenv("FRONT_END_REDIRECT_URL"))
}

func (c Controller) generateTokens(userId int) (accessToken, refreshToken string, err error) {
	accessToken, err = middleware.GenerateJWT(userId)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = c.refreshService.Save(userId)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
