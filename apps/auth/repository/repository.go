package repository

import (
	"context"
	"database/sql"
	"errors"
	"mohhefni/go-online-shop/apps/auth/entity"
	"mohhefni/go-online-shop/infra/errorpkg"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	AddUser(ctx context.Context, model entity.AuthEntity) (id string, err error)
	VerifyAvailableEmail(ctx context.Context, email string) (err error)
	GetUserByEmail(ctx context.Context, email string) (authEntity entity.AuthEntity, err error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) AddUser(ctx context.Context, model entity.AuthEntity) (id string, err error) {
	query := `
		INSERT INTO users (
			public_id, email, password, role, created_at, updated_at
		) VALUES (
			:public_id, :email, :password, :role, :created_at, :updated_at
		) RETURNING public_id
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	err = stmt.GetContext(ctx, &id, model)
	if err != nil {
		return
	}

	return
}

func (r *repository) VerifyAvailableEmail(ctx context.Context, email string) (err error) {
	var count int8
	query := `
		SELECT 
			COUNT(email) 
		FROM users 
		WHERE email = $1
	`

	err = r.db.GetContext(ctx, &count, query, email)
	if err != nil {
		return err
	}

	if count > 0 {
		return errorpkg.ErrEmailAlreadyUsed
	}

	return nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (user entity.AuthEntity, err error) {
	query := `
		SELECT 
			public_id, email, password, role, created_at, updated_at 
		FROM users 
		WHERE email = $1
	`

	err = r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.AuthEntity{}, errorpkg.ErrNotFound
		}
		return entity.AuthEntity{}, err
	}

	return user, nil
}
