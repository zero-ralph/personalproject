package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	form "github.com/zero-ralph/personalporject_users/auth_service/internal/forms"
	"github.com/zero-ralph/personalporject_users/auth_service/internal/services"
	"github.com/zero-ralph/personalporject_users/auth_service/pkg/utils"
)

type AuthenticationAPIInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	HealthCheck(c *gin.Context)
}

type AuthenticationAPIStruct struct {
	service services.AuthenticationServiceInterface
}

func NewAuthenticationAPI(api *gin.RouterGroup, service *services.AuthenticationServiceInterface) {
	controller := &AuthenticationAPIStruct{
		service: *service,
	}

	api.POST("/register", controller.Register)
	api.POST("/login", controller.Login)
	api.GET("/health-check", controller.HealthCheck)
}

func (authenticationAPI *AuthenticationAPIStruct) Register(c *gin.Context) {
	request := &form.RegisterRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		utils.GetErrors(err, c)
		return
	}

	user, err := authenticationAPI.service.Register(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}

func (authenticationAPI *AuthenticationAPIStruct) Login(c *gin.Context) {
	request := &form.AuthenticationRequest{}

	if err := c.ShouldBindJSON(request); err != nil {
		utils.GetErrors(err, c)
		return
	}

	user, err := authenticationAPI.service.Login(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	// Generate Token
	token, err := authenticationAPI.service.GenerateToken(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusGone, gin.H{"errors": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})

}

func (authenticationAPI *AuthenticationAPIStruct) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"detail": "it's working as expected"})
}
