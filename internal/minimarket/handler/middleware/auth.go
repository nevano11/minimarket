package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthValidationService interface {
	IsAuthentificated(ctx context.Context, token string) (bool, error)
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
	token := c.GetHeader("token")

	errorMessage := "your credentials were not found, please log in"

	if len(token) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessage,
		})
		return
	}

	auth, err := x.validationService.IsAuthentificated(c, token)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if auth {
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": errorMessage,
	})
}
