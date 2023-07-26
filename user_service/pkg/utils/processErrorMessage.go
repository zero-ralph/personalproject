package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ProcessErrorMessage(formError validator.FieldError) string {
	switch formError.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return fmt.Sprintf("Should not be less than %v characters", formError.Param())
	}
	return "Test"
}

func GetErrors(err error, c *gin.Context) {
	var validateErrors validator.ValidationErrors
	if errors.As(err, &validateErrors) {
		out := make([]ErrorMessage, len(validateErrors))
		for i, formError := range validateErrors {
			out[i] = ErrorMessage{formError.Field(), ProcessErrorMessage(formError)}
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": out,
		})
	}
}
