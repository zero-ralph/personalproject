package repository

import (
	"errors"

	"github.com/google/uuid"
	form "github.com/zero-ralph/personalporject_users/user_service/internal/forms"
	"github.com/zero-ralph/personalporject_users/user_service/internal/model"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetUserById(userId uuid.UUID) (user *model.User, err error)
	GetUserProfileById(profileId uuid.UUID) (profile *model.Profile, err error)
	SaveProfile(profile *model.Profile, request form.ProfileRequest) (p *model.Profile, err error)
	ChangePassword(user *model.User, password string) (u *model.User, err error)
	TokenSecrets() (jwtHourLifeSpan int, jwtSecret string)
}

type UserRepositoryStruct struct {
	db              *gorm.DB
	jwtHourLifeSpan int
	jwtSecret       string
}

func NewUserRepository(db *gorm.DB, jwtHourLifeSpan int, jwtSecret string) (repository UserRepositoryInterface) {
	repository = &UserRepositoryStruct{
		db:              db,
		jwtHourLifeSpan: jwtHourLifeSpan,
		jwtSecret:       jwtSecret,
	}
	return
}

func (repository *UserRepositoryStruct) GetUserById(userId uuid.UUID) (user *model.User, err error) {
	user = &model.User{}
	if err = repository.db.Model(user).Where("id = ?", userId).Preload("Profile").Take(user).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return
}

func (repository *UserRepositoryStruct) GetUserProfileById(profileId uuid.UUID) (profile *model.Profile, err error) {
	profile = &model.Profile{}
	if err = repository.db.Model(profile).Where("id = ?", profileId).Take(profile).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return
}

func (repository *UserRepositoryStruct) SaveProfile(profile *model.Profile, request form.ProfileRequest) (p *model.Profile, err error) {
	if err = repository.db.Model(profile).Updates(request).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return profile, nil
}

func (repository *UserRepositoryStruct) ChangePassword(user *model.User, password string) (u *model.User, err error) {
	if err = repository.db.Model(user).Update("id", user.ID).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return user, nil
}

func (repository *UserRepositoryStruct) TokenSecrets() (jwtHourLifeSpan int, jwtSecret string) {
	jwtHourLifeSpan, jwtSecret = repository.jwtHourLifeSpan, repository.jwtSecret
	return
}
