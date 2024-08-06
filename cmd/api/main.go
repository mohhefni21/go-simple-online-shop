package main

import (
	"log"
	"mohhefni/go-online-shop/external/database"
	"mohhefni/go-online-shop/internal/config"
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
}
