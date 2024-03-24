package service

import (
	"context"

	"github.com/nevano11/minimarket/internal/minimarket/model"
)

type Storage interface {
	GetPosts(ctx context.Context, filter model.PostFilter, token *string) ([]model.PostAbout, error)
	CreatePost(ctx context.Context, post model.Post, token string) error
}

type AuthValidationService interface {
	IsAuthentificated(ctx context.Context, token string) (bool, error)
}

type Service struct {
	storage       Storage
	authValidator AuthValidationService
}

func NewService(storage Storage, authValidator AuthValidationService) *Service {
	return &Service{
		storage:       storage,
		authValidator: authValidator,
	}
}

func (x *Service) GetPosts(ctx context.Context, filter model.PostFilter, token *string) ([]model.PostAbout, error) {
	return x.storage.GetPosts(ctx, filter, token)
}

func (x *Service) CreatePost(ctx context.Context, post model.Post, token string) error {
	return x.storage.CreatePost(ctx, post, token)
}
