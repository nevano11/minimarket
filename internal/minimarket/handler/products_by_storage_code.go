package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var failedGetQueryParameterMessage = "failed to get storage code from query"

func (h *HttpHandler) getGoodsInStockByCode(ctx *gin.Context) {
	code, hasParameter := ctx.GetQuery("code")
	if !hasParameter {
		logrus.Error("failed to get storage code")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"products": nil,
			"error":    failedGetQueryParameterMessage,
		})
		return
	}

	products, err := h.service.GetProductsByStorageCode(code)
	strError := ""

	status := http.StatusOK
	if err != nil {
		status = http.StatusInternalServerError
		strError = err.Error()
	}

	ctx.JSON(status, gin.H{
		"products": products,
		"error":    strError,
	})
}
