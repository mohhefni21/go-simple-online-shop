package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, db *sqlx.DB) {
	repo := newRepository(db)
	srv := newService(repo)
	usecase := newUsecase(repo, srv)
	handler := newHandler(usecase)

	e.POST("/register", handler.register)
	e.POST("/login", handler.login)
}
