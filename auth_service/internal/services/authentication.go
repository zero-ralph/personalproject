package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	form "github.com/zero-ralph/personalporject_users/auth_service/internal/forms"
	"github.com/zero-ralph/personalporject_users/auth_service/internal/model"
	"github.com/zero-ralph/personalporject_users/auth_service/internal/repository"
)

type AuthenticationServiceInterface interface {
	Register(request *form.RegisterRequest) (user *model.User, err error)
	Login(request *form.AuthenticationRequest) (user *model.User, err error)
	GenerateToken(user *model.User) (token string, err error)
}

type AuthenticationServiceStruct struct {
	repository repository.AuthenticationRepositoryInterface
}

func NewAuthenticationService(repository repository.AuthenticationRepositoryInterface) (service AuthenticationServiceInterface) {
	service = &AuthenticationServiceStruct{
		repository: repository,
	}
	return
}

func (service *AuthenticationServiceStruct) Register(request *form.RegisterRequest) (user *model.User, err error) {
	// Check Password vs Confirm Password
	if request.Password != request.PasswordConfirmation {
		return nil, errors.New(err.Error())
	}

	// Check if Username Exists
	exists, err := service.repository.ExistsByUsername(request.Username)
	if !exists {
		return nil, errors.New(err.Error())
	}

	// Save the form
	user, err = service.repository.Save(&model.User{
		Username: request.Username,
		Password: request.Password,
		Created:  time.Now(),
		Updated:  time.Now(),
		Profile: model.Profile{
			Firstname:  request.Profile.Firstname,
			Middlename: request.Profile.Middlename,
			Lastname:   request.Profile.Lastname,
			Created:    time.Now(),
			Updated:    time.Now(),
		},
	})

	return

}

func (service *AuthenticationServiceStruct) Login(request *form.AuthenticationRequest) (user *model.User, err error) {
	// Get the User by Username
	user, err = service.repository.GetUserByUsername(request.Username)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	// Hash the password and verify it
	if err = user.VerifyPassword(request.Password, user.Password); err != nil {
		return nil, errors.New(err.Error())
	}
	// Return the user
	return
}

func (service *AuthenticationServiceStruct) GenerateToken(user *model.User) (token string, err error) {
	jwtHourLifeSpan, jwtSecret := service.repository.TokenSecrets()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = user.ID
	claims["expiration"] = time.Now().Add(time.Hour * time.Duration(jwtHourLifeSpan)).Unix()

	jwtNewClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtNewClaims.SignedString([]byte(jwtSecret))
	return
}
