package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zero-ralph/personalporject_users/domain_service/internal/model"
	"github.com/zero-ralph/personalporject_users/domain_service/internal/repository"
)

type UserDomainServiceInterface interface {
	GetUserDomains(userId uuid.UUID) (userDomains []model.UserDomainPivot, err error)
	TokenSecrets() (jwtHourLifeSpan int, jwtSecret string)
}

type UserDomainServiceStruct struct {
	repository repository.UserDomainRepositoryInterface
}

func NewUserDomainService(repository repository.UserDomainRepositoryInterface) (service UserDomainServiceInterface) {
	service = &UserDomainServiceStruct{
		repository: repository,
	}
	return
}

func (service *UserDomainServiceStruct) GetUserDomains(userId uuid.UUID) (userDomains []model.UserDomainPivot, err error) {

	if userDomains, err = service.repository.GetUserDomains(userId); err == nil {
		return
	}
	return nil, errors.New(err.Error())
}

func (service *UserDomainServiceStruct) TokenSecrets() (jwtHourLifeSpan int, jwtSecret string) {
	jwtHourLifeSpan, jwtSecret = service.repository.TokenSecrets()
	return
}
