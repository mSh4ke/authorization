package api

import (
	_ "github.com/gorilla/mux"
	"github.com/mSh4ke/authorization/storage"
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
	//users handlers
	a.router.HandleFunc(prefix+"/users/register", a.RegisterUser).Methods("POST")
	a.router.HandleFunc(prefix+"/users/auth", a.Authenticate).Methods("POST")
	a.router.HandleFunc(prefix+"/users/test", a.Authenticate).Methods("GET")

	//admin
	a.router.HandleFunc(prefix+"/admin/createRole", a.CreateRole).Methods("POST")
	a.router.HandleFunc(prefix+"/admin/assignRole", a.CreateRole).Methods("POST")
	a.router.HandleFunc(prefix+"/admin/listRoles", a.CreateRole).Methods("POST")
	a.router.HandleFunc(prefix+"/admin/listPerms/{id}", a.CreateRole).Methods("GET")
	a.router.HandleFunc(prefix+"/admin/addPerm", a.CreateRole).Methods("POST")
	a.router.HandleFunc(prefix+"/admin/removePerm", a.CreateRole).Methods("POST")

	//data handlers
	a.router.HandleFunc(prefix+"{endpoint}", a.RouteHandler("GET")).Methods("GET")
	a.router.HandleFunc(prefix+"{endpoint}", a.RouteHandler("POST")).Methods("POST")
	a.router.HandleFunc(prefix+"{endpoint}", a.RouteHandler("PUT")).Methods("PUT")
	a.router.HandleFunc(prefix+"{endpoint}", a.RouteHandler("DELETE")).Methods("DELETE")
}
func (a *API) configreStorageField() error {
	storage := storage.New(a.config.Storage)

	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
