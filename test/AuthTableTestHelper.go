package test

import (
	"mohhefni/go-online-shop/external/database"
	"mohhefni/go-online-shop/internal/config"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthTableTestHelper struct {
	db *sqlx.DB
}

func NewAuthTableTestHelper() (*AuthTableTestHelper, error) {
	db, err := database.Connection(config.Cfg.Db)
	if err != nil {
		return nil, err
	}

	return &AuthTableTestHelper{
		db: db,
	}, nil
}

func (a *AuthTableTestHelper) AddUser(email string, password string) (public_id string, err error) {
	query := `
		INSERT INTO users (
			public_id, email, password, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING public_id
	`

	err = a.db.QueryRow(query, uuid.New(), email, password, time.Now(), time.Now()).Scan(&public_id)
	if err != nil {
		return
	}

	return
}

func (a *AuthTableTestHelper) CleanTableUser() (err error) {
	query := `
		TRUNCATE users
	`

	_, err = a.db.Exec(query)
	return
}
