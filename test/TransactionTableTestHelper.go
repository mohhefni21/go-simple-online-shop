package test

import (
	"mohhefni/go-online-shop/external/database"
	"mohhefni/go-online-shop/internal/config"

	"github.com/jmoiron/sqlx"
)

type TransactionTableTestHelper struct {
	db *sqlx.DB
}

func NewTransactionTableTestHelper() (*TransactionTableTestHelper, error) {
	db, err := database.Connection(config.Cfg.Db)
	if err != nil {
		return nil, err
	}

	return &TransactionTableTestHelper{
		db: db,
	}, nil
}

// func (a *TransactionTableTestHelper) AddTransaction(email string, password string) (err error) {
// 	query := `
// 		INSERT INTO users (
// 			public_id, email, password, created_at, updated_at
// 		) VALUES (
// 			$1, $2, $3, $4, $5
// 		)
// 	`

// 	_, err = a.db.Exec(query, uuid.New(), email, password, time.Now(), time.Now())

// 	return
// }

func (a *TransactionTableTestHelper) CleanTableTransaction() (err error) {
	query := `
		TRUNCATE transactions
	`

	_, err = a.db.Exec(query)
	return
}
