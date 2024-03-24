package repository

import "context"

func (x *Repository) getUserId(ctx context.Context, query string, params map[string]interface{}) (int, error) {
	ids := []int64{}

	request, err := x.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return 0, err
	}

	if err = request.SelectContext(ctx, &ids, params); err != nil {
		return 0, err
	}

	if len(ids) != 1 {
		return 0, errUserNotFound
	}

	return int(ids[0]), nil
}

func (x *Repository) getUserIdByAuthData(ctx context.Context, login, passwordHash string) (int, error) {
	query := `SELECT id FROM "user" 
          WHERE login = :login AND password_hash = :password_hash`

	return x.getUserId(ctx, query, map[string]interface{}{
		"login":         login,
		"password_hash": passwordHash,
	})
}

func (x *Repository) getUserIdByToken(ctx context.Context, token string) (int, error) {
	query := `SELECT id FROM "user" 
          WHERE token = :token`

	return x.getUserId(ctx, query, map[string]interface{}{
		"token": token,
	})
}
