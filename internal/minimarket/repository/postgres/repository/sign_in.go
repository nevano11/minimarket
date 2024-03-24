package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/nevano11/minimarket/internal/minimarket/model"
	"github.com/nevano11/minimarket/internal/minimarket/repository/postgres/repository/generator"
	log "github.com/sirupsen/logrus"
)

var (
	errUserNotFound = errors.New("user not found")
)

func (x *Repository) SignIn(
	ctx context.Context,
	registrationData model.RegistrationData,
	expirationTime time.Duration,
) (token string, err error) {
	log.Infof("sign in user with login=%s", registrationData.Login)

	userId, err := x.getUserIdByAuthData(ctx, registrationData.Login, registrationData.PasswordHash)
	if err != nil {
		return "", fmt.Errorf("SignIn: failed to get userId: %s", err)
	}

	token, err = generator.GenerateToken(registrationData.Login)
	if err != nil {
		return "", fmt.Errorf("SignIn: failed to generate token: %w", err)
	}

	expiredAt := time.Now().Add(expirationTime)
	request, err := x.db.PrepareNamedContext(ctx, `
		UPDATE "user"
		SET token = :token, expiration_date = :expired_at
		WHERE id = :id
	`)
	if err != nil {
		return "", fmt.Errorf("SignIn: failed to create request: %w", err)
	}

	_, err = request.ExecContext(ctx, map[string]interface{}{
		"token":      token,
		"expired_at": expiredAt,
		"id":         userId,
	})
	if err != nil {
		return "", fmt.Errorf("SignIn: failed to execute request: %w", err)
	}

	return token, nil
}
