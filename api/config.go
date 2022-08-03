package api

import (
	"github.com/mSh4ke/authorization/storage"
)

type Config struct {
	BindAddr      string `toml:"bind_addr"`
	LoggerLevel   string `toml:"logger_level"`
	Storage       *storage.Config
	SecretKey     string    `toml:"secret_key"`
	DefaultRoleId int       `toml:"default_role_id"`
	Servers       *[]string `toml:"servers"`
}

func NewConfig() *Config {
	servers := make([]string, 2)
	return &Config{
		BindAddr:      ":8080",
		LoggerLevel:   "debug",
		Storage:       storage.NewConfig(),
		DefaultRoleId: 1,
		Servers:       &servers,
	}
}
