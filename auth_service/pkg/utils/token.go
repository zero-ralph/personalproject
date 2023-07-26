package utils

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/zero-ralph/personalporject_users/auth_service/internal/repository"
)

func ValidateToken(c *gin.Context, repository repository.AuthenticationRepositoryStruct) (err error) {
	tokenString := ExtractToken(c)
	_, jwtSecret := repository.TokenSecrets()

	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) (token string) {
	token = c.Query("token")
	if token != "" {
		return
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return
}
