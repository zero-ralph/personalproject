package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zero-ralph/personalporject_users/domain_service/internal/model"
	"gorm.io/gorm"
)

type DomainRepositoryInterface interface {
	ExistsByName(name string) (exists bool, err error)
	GetDomainById(domainId uuid.UUID) (domain *model.Domain, err error)
	CreateDomain(domain *model.Domain) (*model.Domain, error)
	TokenSecrets() (jwtHourLifeSpan int, jwtSecret string)
}

type DomainRepositoryStruct struct {
	db              *gorm.DB
	jwtHourLifeSpan int
	jwtSecret       string
}

func NewDomainRepository(db *gorm.DB, jwtHourLifeSpan int, jwtSecret string) (repository DomainRepositoryInterface) {
	repository = &DomainRepositoryStruct{
		db:              db,
		jwtHourLifeSpan: jwtHourLifeSpan,
		jwtSecret:       jwtSecret,
	}
	return
}

func (repository *DomainRepositoryStruct) GetDomainById(domainId uuid.UUID) (domain *model.Domain, err error) {
	d := &model.Domain{}
	if err = repository.db.Model(&model.Domain{}).Where("id = ?", domainId).First(d).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return
}


func (repository *DomainRepositoryStruct) CreateDomain(domain *model.Domain) (*model.Domain, error) {
	if err := repository.db.Save(domain).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return domain, nil
}

func (repository *DomainRepositoryStruct) ExistsByName(name string) (exists bool, err error) {
	domain := &model.Domain{}
	if err = repository.db.Model(domain).Where("name = ?", name).First(domain).Error; err != nil {
		return false, errors.New(err.Error())
	}
	return true, nil
}

func (repository *DomainRepositoryStruct) TokenSecrets() (jwtHourLifeSpan int, jwtSecret string) {
	jwtHourLifeSpan, jwtSecret = repository.jwtHourLifeSpan, repository.jwtSecret
	return
}
