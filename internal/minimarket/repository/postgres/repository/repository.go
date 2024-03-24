package repository

import (
	"github.com/jmoiron/sqlx"
)

// TODO inject from config
const (
	defaultPageSize   = 10
	defaultPageNumber = 1
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(database *sqlx.DB) *Repository {
	return &Repository{db: database}
}
