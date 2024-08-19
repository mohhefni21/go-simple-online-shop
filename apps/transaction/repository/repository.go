package repository

import (
	"context"
	"database/sql"
	"mohhefni/go-online-shop/apps/transaction/entity"
	"mohhefni/go-online-shop/infra/errorpkg"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	DBTransaction
	ProductRepository
	TransactionRepository
}

type DBTransaction interface {
	Begin(ctx context.Context) (tx *sqlx.Tx, err error)
	Roolback(ctx context.Context, tx *sqlx.Tx) (err error)
	Commit(ctx context.Context, tx *sqlx.Tx) (err error)
}

type ProductRepository interface {
	GetDetailProductBySku(ctx context.Context, productSku string) (product entity.ProductJsonEntity, err error)
	UpdateStockProduct(ctx context.Context, tx *sqlx.Tx, product entity.ProductJsonEntity) (err error)
}

type TransactionRepository interface {
	AddTransaction(ctx context.Context, tx *sqlx.Tx, transaction entity.TransactionEntity) (err error)
	GetTransactionByUser(ctx context.Context, publicIdUser string) (transactions []entity.TransactionEntity, err error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
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

func (r *repository) AddTransaction(ctx context.Context, tx *sqlx.Tx, transaction entity.TransactionEntity) (err error) {
	query := `
		INSERT INTO transactions (
			  user_public_id, product_id, product_price, amount, sub_total, platform_fee,
			 grand_total, status, product_snapshot, created_at, updated_at
		) VALUES (
			:user_public_id, :product_id, :product_price, :amount, :sub_total, :platform_fee, 
			:grand_total, :status, :product_snapshot, :created_at, :updated_at
		)
	`
	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, transaction)

	return
}

func (r *repository) UpdateStockProduct(ctx context.Context, tx *sqlx.Tx, product entity.ProductJsonEntity) (err error) {
	query := `
			UPDATE products
			SET stock=:stock
			WHERE id=:id		
	`

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, product)
	return
}

func (r *repository) GetTransactionByUser(ctx context.Context, publicIdUser string) (transactions []entity.TransactionEntity, err error) {
	query := `
		SELECT
			id, sku, name, stock, price, created_at, updated_at
		FROM products
		WHERE sku=$1
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	stmt.ExecContext(ctx, publicIdUser)

	if err != nil {
		if err == sql.ErrNoRows {
			err = errorpkg.ErrNotFound
			return []entity.TransactionEntity{}, nil
		}
		return
	}
	return
}

func (r *repository) Begin(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = r.db.BeginTxx(ctx, &sql.TxOptions{})
	return
}

func (r *repository) Roolback(ctx context.Context, tx *sqlx.Tx) (err error) {
	err = tx.Rollback()
	return
}

func (r *repository) Commit(ctx context.Context, tx *sqlx.Tx) (err error) {
	err = tx.Commit()
	return
}
