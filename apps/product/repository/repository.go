package repository

import (
	"context"
	"database/sql"
	"mohhefni/go-online-shop/apps/product/entity"
	"mohhefni/go-online-shop/infra/errorpkg"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	AddProduct(ctx context.Context, model entity.ProductEntity) (skuProduct string, err error)
	GetAllProducts(ctx context.Context, req entity.ProductPaginationEntity) (products []entity.ProductEntity, err error)
	GetDetailProductBySku(ctx context.Context, sku string) (product entity.ProductEntity, err error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) AddProduct(ctx context.Context, model entity.ProductEntity) (skuProduct string, err error) {
	query := `
		INSERT INTO products (
			sku, name, stock, price, created_at, updated_at
		) VALUES (
			:sku, :name, :stock, :price, :created_at, :updated_at
		) RETURNING sku
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

func (r *repository) GetAllProducts(ctx context.Context, model entity.ProductPaginationEntity) (products []entity.ProductEntity, err error) {
	query := `
		SELECT
			id, sku, name, stock, price, created_at, updated_at
		FROM products
		LIMIT $1 OFFSET $2
	`
	
	err = r.db.SelectContext(ctx, &products, query, model.Size, model.Cursor)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorpkg.ErrNotFound
		}

		return
	}

	return
}

func (r *repository) GetDetailProductBySku(ctx context.Context, sku string) (product entity.ProductEntity, err error) {
	query := `
		SELECT
			id, sku, name, stock, price, created_at, updated_at
		FROM products
		WHERE sku=$1
	`

	err = r.db.GetContext(ctx, &product, query, sku)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errorpkg.ErrNotFound
			return
		}
		return
	}
	return
}
