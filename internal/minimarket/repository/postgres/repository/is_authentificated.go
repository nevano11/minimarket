package repository

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

func (x *Repository) IsAuthentificated(ctx context.Context, token string) (bool, error) {
	log.Infof("is authentificated user with token=%s", token)

	// find user id
	var expirationDates []time.Time
	request, err := x.db.PrepareNamedContext(ctx,
		`SELECT expiration_date FROM "user" 
          		WHERE token = :token`)
	if err != nil {
		return false, err
	}

	err = request.SelectContext(ctx, &expirationDates, map[string]interface{}{
		"token": token,
	})
	if err != nil {
		return false, err
	}

	if len(expirationDates) != 1 {
		return false, errUserNotFound
	}

	return expirationDates[0].Compare(time.Now()) != -1, nil
}
