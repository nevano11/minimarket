package handler

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nevano11/minimarket/internal/minimarket/model"
)

// TODO inject from config
var (
	signInMaxDuration     = time.Second
	signUpMaxDuration     = time.Second
	getPostsMaxDuration   = time.Second
	createPostMaxDuration = time.Second
)

type AuthValidator interface {
	Validate(c *gin.Context)
}

type AuthValidationService interface {
	IsAuthentificated(ctx context.Context, token string) (bool, error)
}

type AuthService interface {
	SignIn(ctx context.Context, registrationData model.RegistrationData) (token string, err error)
	SignUp(ctx context.Context, registrationData model.RegistrationData) error
}

type PostService interface {
	GetPosts(ctx context.Context, filter model.PostFilter, token *string) ([]model.PostAbout, error)
	CreatePost(ctx context.Context, post model.Post, token string) error
}

type HttpHandler struct {
	postService          PostService
	authService          AuthService
	authValidatorService AuthValidationService

	authMiddleware AuthValidator
}

func NewHttpHandler(postService PostService, authService AuthService, authValidatorService AuthValidationService, authValidator AuthValidator) *HttpHandler {
	return &HttpHandler{
		postService:          postService,
		authService:          authService,
		authValidatorService: authValidatorService,

		authMiddleware: authValidator,
	}
}

func (h *HttpHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/sign-up", h.signUp)
	router.POST("/sign-in", h.signIn)

	router.GET("/posts", h.getPosts)

	authorizedGroup := router.Group("/")
	authorizedGroup.Use(h.authMiddleware.Validate)

	authorizedGroup.PUT("/create-post", h.createPost)

	return router
}
