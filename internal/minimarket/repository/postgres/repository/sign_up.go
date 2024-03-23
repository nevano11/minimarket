package repository

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (x *Repository) SignUp(ctx context.Context, login, passwordHash string) error {
	log.Infof("sign up user with login=%s", login)

	tx := x.db.MustBeginTx(ctx, nil)
	tx.PrepareNamedContext(ctx, `
	`)

	if err := tx.Commit(); err != nil {
		return err
	}

	query := fmt.Sprintf("select * from %s(ARRAY [%s])", reserveProducts, strSliceToPgArray(productCodes))

	var isReservingSuccessful bool

	err := x.db.Get(&isReservingSuccessful, query)
	if err != nil {
		return false, err
	}

	log.Infof("reserving products %s result: %t", productCodes, isReservingSuccessful)
	return isReservingSuccessful, nil
}
