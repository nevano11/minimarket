package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nevano11/minimarket/internal/minimarket/model"
	log "github.com/sirupsen/logrus"
)

func (h *HttpHandler) signIn(ctx *gin.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, signInMaxDuration)
	defer cancel()

	// при наличии токена он должен быть обработан и проверен ??
	// пришлось инжектить validatorService. Так он тут не нужен
	token := ctx.GetHeader("token")
	if len(token) > 0 {
		authentificated, err := h.authValidatorService.IsAuthentificated(ctxWithTimeout, token)
		if err != nil {
			log.Errorf("failed to auth with token on sign in")
		} else if authentificated {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			return
		}
	}

	var reqData model.LoginForm

	if err := ctx.BindJSON(&reqData); err != nil {
		log.Errorf("failed to bind Json: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, "invalid data")
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
	token, err := h.authService.SignIn(ctxWithTimeout, data)
	if err != nil {
		log.Errorf("failed to sign in: %s", err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
