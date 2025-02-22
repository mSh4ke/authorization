package main

import (
	"flag"
	"log"

	"github.com/mSh4ke/authorization/api"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	//Скажем, что наше приложение будет на этапе запуска получать путь до конфиг файла из внешнего мира
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {

	//В этот момент происходит инициализация переменной configPath значением
	flag.Parse()
	log.Println("starting server")
	//server instance initialization
	config := api.NewConfig()
	log.Println("config path: ", configPath)
	_, err := toml.DecodeFile(configPath, config) // Десериалзиуете содержимое .toml файла
	if err != nil {
		log.Println("can not find configs file. using default values:", err)
	}

	server := api.New(config)
	defer server.ShutDown()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
