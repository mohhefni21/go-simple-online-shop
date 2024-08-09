package database

import (
	"fmt"
	"mohhefni/go-online-shop/internal/config"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connection(config config.DbConfig) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DbName,
	)

	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	db.SetConnMaxIdleTime(time.Duration(config.ConnectionPool.MaxIdleConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(config.ConnectionPool.MaxLifetimeConnection) * time.Second)
	db.SetMaxOpenConns(int(config.ConnectionPool.MaxOpenConnection))
	db.SetMaxIdleConns(int(config.ConnectionPool.MaxIdleConnection))

	return
}
