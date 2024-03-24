package auth

import (
	"context"
	"time"

	"github.com/nevano11/minimarket/internal/minimarket/model"
)

// TODO inject from config
const expirationTime = time.Hour

type AuthRepository interface {
	IsAuthentificated(ctx context.Context, token string) (bool, error)
	SignIn(ctx context.Context, registrationData model.RegistrationData, expirationTime time.Duration) (token string, err error)
	SignUp(ctx context.Context, registrationData model.RegistrationData) error
}

type Service struct {
	repository AuthRepository
}

func New(repository AuthRepository) *Service {
	return &Service{repository: repository}
}

func (x *Service) IsAuthentificated(ctx context.Context, token string) (bool, error) {
	return x.repository.IsAuthentificated(ctx, token)
}
func (x *Service) SignIn(ctx context.Context, registrationData model.RegistrationData) (token string, err error) {
	return x.repository.SignIn(ctx, registrationData, expirationTime)
}
func (x *Service) SignUp(ctx context.Context, registrationData model.RegistrationData) error {
	return x.repository.SignUp(ctx, registrationData)
}
