package auth

import (
	"mohhefni/go-online-shop/apps/auth/handler"
	"mohhefni/go-online-shop/apps/auth/repository"
	"mohhefni/go-online-shop/apps/auth/service"
	"mohhefni/go-online-shop/apps/auth/usecase"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, db *sqlx.DB) {
	repo := repository.NewRepository(db)
	srv := service.NewService(repo)
	usecase := usecase.NewUsecase(repo, srv)
	handler := handler.NewHandler(usecase)

	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)
}
