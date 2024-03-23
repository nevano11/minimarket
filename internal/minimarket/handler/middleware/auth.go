package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserType int

const (
	UserType_UNKNOWN = iota
	UserType_CLIENT
)

type AuthValidationService interface {
	ValidateAuth() (UserType, error)
}

type Auth struct {
	validationService AuthValidationService
}

func New(validationService AuthValidationService) *Auth {
	return &Auth{
		validationService,
	}
}

func (x *Auth) Validate(c *gin.Context) {
	auth, err := x.validationService.ValidateAuth()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if auth == UserType_UNKNOWN {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "your credentials were not found, please log in",
		})
		return
	}

	c.Next()
}
