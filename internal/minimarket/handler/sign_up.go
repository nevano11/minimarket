package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nevano11/minimarket/internal/minimarket/model"
	log "github.com/sirupsen/logrus"
)

func (h *HttpHandler) signUp(ctx *gin.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, signUpMaxDuration)
	defer cancel()

	var reqData model.LoginForm

	if err := ctx.BindJSON(&reqData); err != nil {
		log.Errorf("failed to bind Json: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(reqData); err != nil {
		log.Errorf("validating of login form failed: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "the login and password are not secure enough",
		})
		return
	}

	data := reqData.ToRegistrationData()
	if err := h.authService.SignUp(ctxWithTimeout, data); err != nil {
		log.Errorf("failed to sign up: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to sign up. May be user with that login is already exists",
		})
		return
	}

	ctx.JSON(http.StatusOK, reqData)
}
