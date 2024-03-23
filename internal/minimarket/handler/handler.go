package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nevano11/minimarket/internal/minimarket/service"
)

type HttpHandler struct {
	service *service.Service
}

func NewHttpHandler(service *service.Service) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}

func (h *HttpHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/reserve", h.reserveProducts)
	router.POST("/free", h.freeReserved)

	router.GET("/goods-in-stock", h.getGoodsInStockByCode)

	return router
}
