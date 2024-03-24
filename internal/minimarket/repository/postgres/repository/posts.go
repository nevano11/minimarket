package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/nevano11/minimarket/internal/minimarket/model"
	log "github.com/sirupsen/logrus"
)

func (x *Repository) GetPosts(ctx context.Context, filter model.PostFilter, token *string) ([]model.PostAbout, error) {
	log.Infof("trying to get posts with filter %s", filter.String())

	if filter.PageNumber == nil || *filter.PageNumber < defaultPageNumber {
		tmp := defaultPageNumber
		filter.PageNumber = &tmp
	}
	if filter.PageSize == nil || *filter.PageNumber <= 0 {
		tmp := defaultPageSize
		filter.PageSize = &tmp
	}

	var id *int
	if token != nil {
		userId, err := x.getUserIdByToken(ctx, *token)
		if err != nil {
			log.Warnf("failed to get userId: %s", err.Error())
		} else {
			id = &userId
		}
	}

	queryBuilder := strings.Builder{}
	queryBuilder.WriteString(`SELECT name, description, img "image", price, u.login login, user_id = :uid as is_current_user 
			FROM post p
			LEFT JOIN "user" u ON u.id = p.user_id 
          	WHERE true`)

	if filter.MinPrice != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND price >= %d", *filter.MinPrice))
	}
	if filter.MinPrice != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND price <= %d", *filter.MaxPrice))
	}
	if filter.PriceOrder != model.None || filter.DateOrder != model.None {
		orderByAdded := false
		if filter.PriceOrder != model.None {
			queryBuilder.WriteString(" ORDER BY")
			orderByAdded = true

			if filter.PriceOrder == model.Asc {
				queryBuilder.WriteString(" price ASC")
			} else {
				queryBuilder.WriteString(" price DESC")
			}
		}

		if filter.DateOrder != model.None {
			if orderByAdded {
				queryBuilder.WriteString(", ")
			} else {
				queryBuilder.WriteString(" ORDER BY")
			}

			if filter.DateOrder == model.Asc {
				queryBuilder.WriteString(" created_at ASC")
			} else {
				queryBuilder.WriteString(" created_at DESC")
			}
		}
	}

	queryBuilder.WriteString(" LIMIT :limit OFFSET :offset;")

	log.Infof("execute select: %s", queryBuilder.String())

	request, err := x.db.PrepareNamedContext(ctx, queryBuilder.String())
	if err != nil {
		return nil, err
	}

	posts := []model.PostAbout{}
	err = request.SelectContext(ctx, &posts, map[string]interface{}{
		"uid":    id,
		"limit":  filter.PageSize,
		"offset": (*filter.PageNumber - 1) * *filter.PageSize,
	})
	if err != nil {
		return nil, err
	}

	return posts, nil
}
