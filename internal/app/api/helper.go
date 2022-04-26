package api

import (
	"GitHab/Autorization/storage"
	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	prefix = "/api/v2"
)

func (a *API) configreLoggerField() error {

	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configreRouterField() {

	a.router.HandleFunc(prefix+"/users/registrate", a.RegistratedUser).Methods("POST")
	a.router.HandleFunc(prefix+"/users/auth", a.PostToAuth).Methods("POST")
	a.router.HandleFunc("/about", a.Tok).Methods("GET")
	a.router.HandleFunc("/images", a.Token).Methods("GET")
}
func (a *API) configreStorageField() error {
	storage := storage.New(a.config.Storage)

	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
