package auth

import (
	"mohhefni/go-online-shop/apps/auth/handler"
	"mohhefni/go-online-shop/apps/auth/repository"
	"mohhefni/go-online-shop/apps/auth/usecase"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, db *sqlx.DB) {
	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo)
	handler := handler.NewHandler(usecase)

	g := e.Group("auth")
	g.POST("/register", handler.PostRegisterHandler)
	g.POST("/login", handler.PostLoginHandler)
}
