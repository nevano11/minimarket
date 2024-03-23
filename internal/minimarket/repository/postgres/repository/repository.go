package repository

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

const (
	selectProductsByStorageCode = "select_products_by_storage"
	reserveProducts             = "reserve_products"
	freeProducts                = "free_products"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(database *sqlx.DB) *Repository {
	return &Repository{db: database}
}

func strSliceToPgArray(sl []string) string {
	pgArray := strings.Builder{}
	for i, code := range sl {
		pgArray.WriteString("'" + code + "'")
		if i != len(sl)-1 {
			pgArray.WriteString(", ")
		}
	}
	return pgArray.String()
}
