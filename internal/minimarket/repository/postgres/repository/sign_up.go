package repository

import (
	"context"

	"github.com/nevano11/minimarket/internal/minimarket/model"
	log "github.com/sirupsen/logrus"
)

func (x *Repository) SignUp(ctx context.Context, registrationData model.RegistrationData) error {
	log.Infof("sign up user with login=%s", registrationData.Login)

	ns, err := x.db.PrepareNamedContext(ctx, `
		INSERT INTO "user" (login, password_hash)
		VALUES (:login, :password_hash)
	`)
	if err != nil {
		return err
	}

	if _, err := ns.ExecContext(ctx, registrationData); err != nil {
		return err
	}

	return nil
}
