package transaction

import (
	"mohhefni/go-online-shop/apps/auth/entity"
	"mohhefni/go-online-shop/apps/transaction/handler"
	"mohhefni/go-online-shop/apps/transaction/repository"
	"mohhefni/go-online-shop/apps/transaction/usecase"
	"mohhefni/go-online-shop/infra/middleware"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, db *sqlx.DB) {
	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo)
	handler := handler.NewHandler(usecase)

	g := e.Group("transactions", middleware.CheckAuth, middleware.CheckRole([]string{string(entity.ROLE_USER)}))
	g.POST("/checkout", handler.PostTransactionHandler)
	g.GET("/history", handler.GetTransactionByUserHandler)
}
