package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nevano11/minimarket/internal/minimarket/service"
)

type AuthValidator interface {
	Validate(c *gin.Context)
}

type HttpHandler struct {
	service        *service.Service
	authMiddleware AuthValidator
}

func NewHttpHandler(service *service.Service, authValidator AuthValidator) *HttpHandler {
	return &HttpHandler{
		service: service,

		authMiddleware: authValidator,
	}
}

func (h *HttpHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/sign-up", h.signUp)
	router.POST("/sign-in", h.signIn)

	router.GET("/announcements", h.announcements)

	authorizedGroup := router.Group("/")
	authorizedGroup.Use(h.authMiddleware.Validate)

	authorizedGroup.PUT("/create-announcement", h.createAnnouncement)

	return router
}
