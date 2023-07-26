package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	form "github.com/zero-ralph/personalporject_users/domain_service/internal/forms"
	"github.com/zero-ralph/personalporject_users/domain_service/internal/model"
	"github.com/zero-ralph/personalporject_users/domain_service/internal/repository"
)

type DomainServiceInterface interface {
	GetDomains() (domains []*model.Domain, err error)
	CreateDomain(request *form.DomainRequest) (domain *model.Domain, err error)
	AttachDomainToUser(domainId uuid.UUID, userId uuid.UUID) (isAttached bool, err error)
	TokenSecrets() (jwtHourLifeSpan int, jwtSecret string)
}

type DomainServiceStruct struct {
	repository           repository.DomainRepositoryInterface
	userDomainRepository repository.UserDomainRepositoryInterface
}

func NewDomainService(repository repository.DomainRepositoryInterface, userDomainRepository repository.UserDomainRepositoryInterface) (service DomainServiceInterface) {
	service = &DomainServiceStruct{
		repository:           repository,
		userDomainRepository: userDomainRepository,
	}
	return
}

func (service *DomainServiceStruct) GetDomains() (domain []*model.Domain, err error) {
	return
}

func (service *DomainServiceStruct) CreateDomain(request *form.DomainRequest) (domain *model.Domain, err error) {
	// Get request name and check if the domain exists
	exists, _ := service.repository.ExistsByName(request.Name)
	if exists {
		return nil, errors.New("domain is exists")
	}
	// If not proceed to saving
	domain, err = service.repository.CreateDomain(&model.Domain{
		Name:    request.Name,
		Created: time.Now(),
		Updated: time.Now(),
	})
	return
}

func (service *DomainServiceStruct) AttachDomainToUser(domainId uuid.UUID, userId uuid.UUID) (isAttached bool, err error) {
	// Negate the attachment first
	isAttached = false
	// Get the domain
	_, err = service.repository.GetDomainById(domainId)
	if err != nil {
		return isAttached, errors.New(err.Error())
	}
	// Add a record to UserDomainPivot table
	fmt.Println(domainId, userId)
	_, err = service.userDomainRepository.AttachDomainToUser(&model.UserDomainPivot{
		DomainId: domainId,
		UserId:   userId,
		Created:  time.Now(),
		Updated:  time.Now(),
	})
	if err != nil {
		return isAttached, errors.New(err.Error())
	}
	isAttached = true
	return
}

func (service *DomainServiceStruct) TokenSecrets() (jwtHourLifeSpan int, jwtSecret string) {
	jwtHourLifeSpan, jwtSecret = service.repository.TokenSecrets()
	return
}
