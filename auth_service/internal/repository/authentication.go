package repository

import (
	"errors"

	"github.com/zero-ralph/personalporject_users/auth_service/internal/model"
	"gorm.io/gorm"
)

type AuthenticationRepositoryInterface interface {
	Save(user *model.User) (*model.User, error)
	ExistsByUsername(username string) (exists bool, err error)
	GetUserByUsername(username string) (user *model.User, err error)
	TokenSecrets() (jwtHourLifeSpan int, jwtSecret string)
}

type AuthenticationRepositoryStruct struct {
	db              *gorm.DB
	jwtHourLifeSpan int
	jwtSecret       string
}

func NewAuthenticationRepository(db *gorm.DB, jwtHourLifeSpan int, jwtSecret string) (repository AuthenticationRepositoryInterface) {
	repository = &AuthenticationRepositoryStruct{
		db:              db,
		jwtHourLifeSpan: jwtHourLifeSpan,
		jwtSecret:       jwtSecret,
	}
	return
}

func (repository *AuthenticationRepositoryStruct) Save(user *model.User) (*model.User, error) {
	if err := repository.db.Save(user).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return user, nil
}

func (repository *AuthenticationRepositoryStruct) ExistsByUsername(username string) (exists bool, err error) {
	user := &model.User{}
	if err := repository.db.Model(user).Select("count(*) > 0").Where("username = ?", username).Find(&exists).Error; err != nil {
		return false, errors.New(err.Error())
	}
	return true, nil
}

func (repository *AuthenticationRepositoryStruct) GetUserByUsername(username string) (user *model.User, err error) {
	user = &model.User{}
	if err := repository.db.Model(user).Where("username = ?", username).Preload("Profile").Take(user).Error; err != nil {
		return nil, err
	}
	return
}

func (repository *AuthenticationRepositoryStruct) TokenSecrets() (jwtHourLifeSpan int, jwtSecret string) {
	jwtHourLifeSpan, jwtSecret = repository.jwtHourLifeSpan, repository.jwtSecret
	return
}
