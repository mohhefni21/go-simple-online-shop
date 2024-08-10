package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
	Db  DbConfig  `yaml:"db"`
}

type AppConfig struct {
	Name      string          `yaml:"name"`
	Port      string          `yaml:"port"`
	Encrytion EncrytionConfig `yaml:"encrytion"`
}

type EncrytionConfig struct {
	Salt      uint8  `yaml:"salt"`
	JWTSecret string `yaml:"jwt_secret"`
}

type DbConfig struct {
	Host           string                 `yaml:"host"`
	Port           string                 `yaml:"port"`
	Username       string                 `yaml:"username"`
	Password       string                 `yaml:"password"`
	DbName         string                 `yaml:"dbname"`
	ConnectionPool DbConnectionPoolConfig `yaml:"connection_pool"`
}

type DbConnectionPoolConfig struct {
	MaxIdleConnection     uint8 `yaml:"max_idle_connection"`
	MaxOpenConnection     uint8 `yaml:"max_open_connection"`
	MaxLifetimeConnection uint8 `yaml:"max_lifetime_connection"`
	MaxIdletimeConnection uint8 `yaml:"max_idletime_connection"`
}

var Cfg Config

func LoadConfig(filename string) (err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	return yaml.Unmarshal(data, &Cfg)
}
