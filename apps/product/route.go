package product

import (
	"mohhefni/go-online-shop/apps/product/handler"
	"mohhefni/go-online-shop/apps/product/repository"
	"mohhefni/go-online-shop/apps/product/usecase"
	"mohhefni/go-online-shop/infra/middleware"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, db *sqlx.DB) {
	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repo)
	handler := handler.NewHandler(usecase)

	g := e.Group("products", middleware.CheckAuth)
	g.POST("", handler.PostProductHandler)
	g.GET("", handler.GetAllProductsHandler)
	g.GET("/sku/:sku", handler.GetDetailProductHandler)
}
