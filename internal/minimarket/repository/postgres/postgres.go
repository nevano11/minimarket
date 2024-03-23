package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nevano11/minimarket/pkg/config"
)

func NewPostgresDb(config config.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			config.Host,
			config.Port,
			config.Username,
			config.DbName,
			config.Password,
			config.SslMode,
		))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
