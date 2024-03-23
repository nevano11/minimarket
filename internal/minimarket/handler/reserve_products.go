package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nevano11/minimarket/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func (h *HttpHandler) reserveProducts(ctx *gin.Context) {
	var productCodes []string

	if err := ctx.BindJSON(&productCodes); err != nil {
		log.Errorf("failed to bind Json: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"reserved": false,
			"error":    "invalid data. expected []string",
		})
		return
	}

	if !utils.IsUnique(productCodes) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"reserved": false,
			"error":    "codes must be unique",
		})
		return
	}

	resultOfReserving, err := h.service.ReserveProducts(productCodes)
	strError := ""
	status := http.StatusOK

	if err != nil {
		status = http.StatusInternalServerError
		strError = err.Error()
	}

	ctx.JSON(status, gin.H{
		"reserved": resultOfReserving,
		"error":    strError,
	})
}
