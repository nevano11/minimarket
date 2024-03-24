package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nevano11/minimarket/internal/minimarket/model"
	log "github.com/sirupsen/logrus"
)

func (h *HttpHandler) createPost(ctx *gin.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, createPostMaxDuration)
	defer cancel()

	token := ctx.GetHeader("token")
	if len(token) == 0 {
		log.Warnf("failed to parse token from header")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "authorization token not found",
		})
		return
	}

	var reqData model.Post

	if err := ctx.BindJSON(&reqData); err != nil {
		log.Errorf("failed to bind Json: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(reqData); err != nil {
		log.Errorf("validating of login form failed: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "errors on data",
		})
		return
	}

	if err := h.postService.CreatePost(ctxWithTimeout, reqData, token); err != nil {
		log.Errorf("validating of login form failed: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create post",
		})
		return
	}

	ctx.JSON(http.StatusOK, reqData)
}
