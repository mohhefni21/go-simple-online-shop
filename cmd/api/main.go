package main

import (
	"fmt"
	"log"
	"mohhefni/go-online-shop/apps/auth"
	"mohhefni/go-online-shop/external/database"
	"mohhefni/go-online-shop/internal/config"

	"github.com/labstack/echo/v4"
)

func main() {
	filename := "cmd/api/config.yaml"

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

	auth.Init(e, db)

	addr := fmt.Sprint("127.0.0.1", config.Cfg.App.Port)
	fmt.Printf("starting web server at %s", addr)

	err = e.Start(addr)
	if err != nil {
		panic(err)
	}
}
