package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zero-ralph/personalporject_users/domain_service/internal/model"
	"gorm.io/gorm"
)

type UserDomainRepositoryInterface interface {
	AttachDomainToUser(userDomain *model.UserDomainPivot) (*model.UserDomainPivot, error)
	GetUserDomains(userId uuid.UUID) (userDomains []model.UserDomainPivot, err error)
	TokenSecrets() (jwtHourLifeSpan int, jwtSecret string)
}

type UserDomainRepositoryStruct struct {
	db              *gorm.DB
	jwtHourLifeSpan int
	jwtSecret       string
}

func NewUserDomainRepository(db *gorm.DB, jwtHourLifeSpan int, jwtSecret string) (repository UserDomainRepositoryInterface) {
	repository = &UserDomainRepositoryStruct{
		db:              db,
		jwtHourLifeSpan: jwtHourLifeSpan,
		jwtSecret:       jwtSecret,
	}
	return
}

func (userDomainRepository *UserDomainRepositoryStruct) AttachDomainToUser(userDomain *model.UserDomainPivot) (*model.UserDomainPivot, error) {
	if err := userDomainRepository.db.Save(userDomain).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return userDomain, nil
}

func (userDomainRepository *UserDomainRepositoryStruct) GetUserDomains(userId uuid.UUID) (uDs []model.UserDomainPivot, err error) {
	var userDomains []model.UserDomainPivot
	if err = userDomainRepository.db.Model(&model.UserDomainPivot{}).Where("user_id = ?", userId).Find(&userDomains).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	uDs = userDomains
	return
}

func (repository *UserDomainRepositoryStruct) TokenSecrets() (jwtHourLifeSpan int, jwtSecret string) {
	jwtHourLifeSpan, jwtSecret = repository.jwtHourLifeSpan, repository.jwtSecret
	return
}
