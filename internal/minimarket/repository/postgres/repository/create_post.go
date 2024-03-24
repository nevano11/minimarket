package repository

import (
	"context"
	"time"

	"github.com/nevano11/minimarket/internal/minimarket/model"
	log "github.com/sirupsen/logrus"
)

func (x *Repository) CreatePost(ctx context.Context, post model.Post, token string) error {
	log.Infof("create post %s by user with token=%s", post.String(), token)

	id, err := x.getUserIdByToken(ctx, token)
	if err != nil {
		return err
	}

	ns, err := x.db.PrepareNamedContext(ctx, `
		INSERT INTO post (user_id, name, price, img, created_at)
		VALUES (:uid, :name, :price, :img, :created_at)
	`)
	if err != nil {
		return err
	}

	if _, err = ns.ExecContext(ctx, map[string]interface{}{
		"uid":        id,
		"name":       post.Name,
		"price":      post.Price,
		"img":        post.Image,
		"created_at": time.Now(),
	}); err != nil {
		return err
	}

	return nil
}
