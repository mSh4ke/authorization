package api

import (
	"GitHab/Autorization/storage"
)

type Config struct {
	//Port
	BindAddr string `toml:"bind_addr"`
	//Logger Level
	LoggerLevel string `toml:"logger_level"`
	//Store configs
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8083",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
