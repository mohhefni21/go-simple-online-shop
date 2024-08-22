package test

import (
	"mohhefni/go-online-shop/external/database"
	"mohhefni/go-online-shop/internal/config"
	"time"

	"github.com/jmoiron/sqlx"
)

type ProductTableTestHelper struct {
	db *sqlx.DB
}

func NewProductTableTestHelper() (*ProductTableTestHelper, error) {
	db, err := database.Connection(config.Cfg.Db)
	if err != nil {
		return nil, err
	}

	return &ProductTableTestHelper{
		db: db,
	}, nil
}

func (a *ProductTableTestHelper) AddProduct(sku string, name string, stock int16, price int) (err error) {
	query := `
		INSERT INTO products (
			sku, name, stock, price, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`

	_, err = a.db.Exec(query, sku, name, stock, price, time.Now(), time.Now())

	return
}

func (p *ProductTableTestHelper) CleanTableProduct() (err error) {
	query := `
		TRUNCATE products
	`

	_, err = p.db.Exec(query)
	return
}
