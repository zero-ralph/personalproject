package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	form "github.com/zero-ralph/personalporject_users/user_service/internal/forms"
	"github.com/zero-ralph/personalporject_users/user_service/internal/middleware"
	"github.com/zero-ralph/personalporject_users/user_service/internal/services"
)

type UserAPIInterface interface {
	GetUserById(c *gin.Context)
	GetUserProfileById(c *gin.Context)
	SaveProfile(c *gin.Context)
	ChangePassword(c *gin.Context)
	HealthCheck(c *gin.Context)
}

type UserAPIStruct struct {
	service services.UserServiceInterface
}

func NewUserAPI(api *gin.RouterGroup, service *services.UserServiceInterface) {
	controller := &UserAPIStruct{
		service: *service,
	}

	api.GET("/health-check", controller.HealthCheck)
	api.Use(middleware.JWTAuthMiddleware(service))
	api.GET("/user/:id", controller.GetUserById)
	api.GET("/user/:id/:profileId", controller.GetUserProfileById)
	api.PUT("/user/:id/:profileId/update", controller.SaveProfile)
	api.POST("/user/:id/change-password", controller.ChangePassword)

}

func (userApi *UserAPIStruct) GetUserById(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userApi.service.GetUserById(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (userApi *UserAPIStruct) GetUserProfileById(c *gin.Context) {
	profileId, err := uuid.Parse(c.Param("profileId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := userApi.service.GetUserProfileById(profileId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

func (userApi *UserAPIStruct) SaveProfile(c *gin.Context) {
	profileId, err := uuid.Parse(c.Param("profileId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	request := &form.ProfileRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := userApi.service.SaveProfile(request, profileId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"profile": profile})

}

func (userApi *UserAPIStruct) ChangePassword(c *gin.Context) {
	// Get the URL id Parameters
	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Bind the request Form
	request := &form.ChangePasswordRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pass to Service
	_, err = userApi.service.ChangePassword(request, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"detail": "Successfully changed the password"})

}

func (userApi *UserAPIStruct) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"detail": "It's working as expected"})
}
