package repository

import (
	"context"
	"database/sql"
	"mohhefni/go-online-shop/apps/transaction/entity"
	"mohhefni/go-online-shop/infra/errorpkg"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	GetDetailProductBySku(ctx context.Context, sku string) (product entity.ProductJsonEntity, err error)
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) GetDetailProductBySku(ctx context.Context, sku string) (product entity.ProductJsonEntity, err error) {
	query := `
		SELECT
			id, sku, name, stock, price
		FROM products
		WHERE sku=$1
	`

	err = p.db.GetContext(ctx, &product, query, sku)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errorpkg.ErrNotFound
			return
		}
		return
	}
	return
}
