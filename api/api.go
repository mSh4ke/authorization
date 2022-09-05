package api

import (
	"net/http"

	"github.com/mSh4ke/authorization/storage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	Config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		Config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port:", api.Config.BindAddr)
	api.logger.Info("Default user role id: ", api.Config.DefaultRoleId)
	api.logger.Info("Secret key: ", api.Config.SecretKey)
	api.configureRouterField()
	if err := api.configureStorageField(); err != nil {
		return err
	}
	if err := api.storage.UserRepository.InitAdmin(); err != nil {
		return err
	}
	return http.ListenAndServe(api.Config.BindAddr, api.router)
}

func (api *API) ShutDown() {
	api.storage.Close()
	return
}

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}
