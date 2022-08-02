package api

import (
	"github.com/mSh4ke/authorization/storage"
)

type Config struct {
	BindAddr      string `toml:"bind_addr"`
	LoggerLevel   string `toml:"logger_level"`
	Storage       *storage.Config
	SecretKey     string `toml:"secret_key"`
	DefaultRoleId int    `toml:"default_role_id"`
	DataServerUrl string `toml:"data_server_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:      ":8083",
		LoggerLevel:   "debug",
		Storage:       storage.NewConfig(),
		DefaultRoleId: 1,
	}
}
