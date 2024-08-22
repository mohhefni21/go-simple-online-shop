package main

import (
	"fmt"
	"log"
	"mohhefni/go-online-shop/apps/auth"
	"mohhefni/go-online-shop/apps/product"
	"mohhefni/go-online-shop/apps/transaction"
	"mohhefni/go-online-shop/external/database"
	inframiddleware "mohhefni/go-online-shop/infra/middleware"
	"mohhefni/go-online-shop/internal/config"
	"mohhefni/go-online-shop/utility"
	
	_ "mohhefni/go-online-shop/docs"
	
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Go online shop api
// @version 1.0
// @description This is a simple API for an online shop with go, providing functionalities like product management, order processing, and user management.

// @host localhost:9000

func main() {
	filename := "config.yaml"

	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.Connection(config.Cfg.Db)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	e := echo.New()

	e.Use(inframiddleware.Logging)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
      AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Authorization"},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	auth.Init(e, db)
	product.Init(e, db)
	transaction.Init(e, db)

	addr := fmt.Sprint("127.0.0.1", config.Cfg.App.Port)
	fmt.Printf("starting web server at %s \n", addr)
	utility.MakeLogEntry(nil).Warning("application started without ssl/tls enabled")
	err = e.Start(addr)

	if err != nil {
		utility.MakeLogEntry(nil).Panic(err)
	}
}
