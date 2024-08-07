package auth

import (
	"mohhefni/go-online-shop/external/database"
	"mohhefni/go-online-shop/internal/config"
)

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)

	if err != nil {
		panic(err)
	}

	db, err := database.Connection(config.Cfg.Db)
	if err != nil {
		panic(err)
	}

	repository := newRepository(db)
	service := newService(repository)

}
