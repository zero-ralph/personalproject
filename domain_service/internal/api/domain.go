package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	form "github.com/zero-ralph/personalporject_users/domain_service/internal/forms"
	"github.com/zero-ralph/personalporject_users/domain_service/internal/middleware"
	"github.com/zero-ralph/personalporject_users/domain_service/internal/services"
	"github.com/zero-ralph/personalporject_users/domain_service/pkg/utils"
)

type DomainAPIInterface interface {
	HealthCheck(c *gin.Context)
	CreateDomain(c *gin.Context)
	AttachDomainTo(c *gin.Context)
	GetDomains(c *gin.Context)
}

type DomainAPIStruct struct {
	service services.DomainServiceInterface
}

func NewDomainAPI(api *gin.RouterGroup, service *services.DomainServiceInterface) {
	controller := &DomainAPIStruct{
		service: *service,
	}

	api.GET("/health-check", controller.HealthCheck)
	api.Use(middleware.JWTAuthMiddleware(service))
	api.POST("/domain/", controller.CreateDomain)
	api.POST("/domain/:domainId/attached_to/:userId", controller.AttachDomainToUser)
}

func (domainAPI *DomainAPIStruct) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"detail": "It's working as expected"})
}

func (domainAPI *DomainAPIStruct) CreateDomain(c *gin.Context) {
	request := &form.DomainRequest{}

	if err := c.ShouldBindJSON(request); err != nil {
		utils.GetErrors(err, c)
		return
	}
	domain, err := domainAPI.service.CreateDomain(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": domain})
}

func (domainAPI *DomainAPIStruct) AttachDomainToUser(c *gin.Context) {
	domainId, err := uuid.Parse(c.Param("domainId"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	isAttached, err := domainAPI.service.AttachDomainToUser(domainId, userId)

	if !isAttached {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"detail": "Domains attached"})
}

func (domainAPI *DomainAPIStruct) GetDomains(c *gin.Context) {

}
