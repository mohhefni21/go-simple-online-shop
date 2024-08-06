package database

import (
	"mohhefni/go-online-shop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)

	if err != nil {
		panic(err)
	}
}

func TestConnectionPostgres(t *testing.T) {
	t.Run("should successfully connect to the database", func(t *testing.T) {
		db, err := Connection(config.Cfg.Db)

		require.Nil(t, err)
		require.NotNil(t, db)
	})
}
