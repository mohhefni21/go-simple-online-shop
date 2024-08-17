package repository

import (
	"context"
	"database/sql"
	"mohhefni/go-online-shop/apps/transaction/entity"
	"mohhefni/go-online-shop/infra/errorpkg"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	ProductRepository
	TransactionRepository
}

type ProductRepository interface {
	GetDetailProductBySku(ctx context.Context, productSku string) (product entity.ProductJsonEntity, err error)
	// UpdateStockProduct(ctx context.Context, productId int, newStock uint8) (err error)
}

type TransactionRepository interface {
	AddTransaction(ctx context.Context, transaction entity.TransactionEntity) (err error)
}

type repository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetDetailProductBySku(ctx context.Context, productSku string) (product entity.ProductJsonEntity, err error) {
	query := `
		SELECT
			id, sku, name, stock, price
		FROM products
		WHERE sku=$1
	`

	err = r.db.GetContext(ctx, &product, query, productSku)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errorpkg.ErrNotFound
			return
		}
		return
	}
	return
}

func (r *repository) AddTransaction(ctx context.Context, transaction entity.TransactionEntity) (err error) {
	query := `
		INSERT INTO products (
			 id, email, product_id, product_price, amount, sub_total, platform_fee,
			 grand_total, status, product_snapshot, created_at, updated_at
		) VALUES (
			:id, :email, :product_id, :product_price, :amount, :sub_total, :platform_fee,
:			 :grand_total, :status, :product_snapshot, :created_at, :updated_at
		)
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	err = stmt.GetContext(ctx, &skuProduct, model)
	if err != nil {
		return
	}

	return
}
