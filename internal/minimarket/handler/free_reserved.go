package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nevano11/minimarket/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func (h *HttpHandler) freeReserved(ctx *gin.Context) {
	var productCodes []string

	if err := ctx.BindJSON(&productCodes); err != nil {
		log.Errorf("failed to bind Json: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid data. expected []string",
		})
		return
	}

	if !utils.IsUnique(productCodes) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "codes must be unique",
		})
		return
	}

	isFreeSuccess, err := h.service.FreeProducts(productCodes)
	strError := ""
	status := http.StatusOK

	if err != nil {
		status = http.StatusInternalServerError
		strError = err.Error()
	}

	ctx.JSON(status, gin.H{
		"success": isFreeSuccess,
		"error":   strError,
	})
}
