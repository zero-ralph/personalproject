package services

import (
	"errors"

	"github.com/google/uuid"
	form "github.com/zero-ralph/personalporject_users/user_service/internal/forms"
	"github.com/zero-ralph/personalporject_users/user_service/internal/model"
	"github.com/zero-ralph/personalporject_users/user_service/internal/repository"
)

type UserServiceInterface interface {
	GetUserById(userId uuid.UUID) (user *model.User, err error)
	GetUserProfileById(profileId uuid.UUID) (profile *model.Profile, err error)
	SaveProfile(request *form.ProfileRequest, profileId uuid.UUID) (profile *model.Profile, err error)
	ChangePassword(request *form.ChangePasswordRequest, userId uuid.UUID) (user *model.User, err error)
	TokenSecrets() (jwtHourLifeSpan int, jwtSecret string)
}

type UserServiceStruct struct {
	repository repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface) (service UserServiceInterface) {
	service = &UserServiceStruct{
		repository: repository,
	}
	return
}

func (service *UserServiceStruct) GetUserById(userId uuid.UUID) (user *model.User, err error) {
	user, err = service.repository.GetUserById(userId)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return
}

func (service *UserServiceStruct) GetUserProfileById(profileId uuid.UUID) (profile *model.Profile, err error) {
	profile, err = service.repository.GetUserProfileById(profileId)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return
}

func (service *UserServiceStruct) SaveProfile(request *form.ProfileRequest, profileId uuid.UUID) (profile *model.Profile, err error) {
	// Get the Profile by Id
	profile, err = service.repository.GetUserProfileById(profileId)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// Then Save
	profile, err = service.repository.SaveProfile(profile, *request)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return
}

func (service *UserServiceStruct) ChangePassword(request *form.ChangePasswordRequest, userId uuid.UUID) (user *model.User, err error) {

	// Validate the request's new password vs new password confirmation
	if request.NewPassword != request.NewPasswordConfirmation {
		return nil, errors.New(err.Error())
	}

	// Look up in Database if that id provided exists or has a record
	user, err = service.repository.GetUserById(userId)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// Validate the Old Password vs the Current User Password
	if err = user.VerifyOldPassword(request.OldPassword); err != nil {
		return nil, errors.New(err.Error())
	}

	// Update the User's Password
	user.Password = request.NewPassword
	user, err = service.repository.ChangePassword(user, request.NewPassword)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return
}

func (service *UserServiceStruct) TokenSecrets() (jwtHourLifeSpan int, jwtSecret string) {
	jwtHourLifeSpan, jwtSecret = service.repository.TokenSecrets()
	return
}
