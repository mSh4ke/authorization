package api

import (
	_ "github.com/gorilla/mux"
	"github.com/mSh4ke/authorization/storage"
	"github.com/sirupsen/logrus"
)

var (
	prefix = "/api/v2"
)

func (a *API) configureLoggerField() error {

	log_level, err := logrus.ParseLevel(a.Config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}
func (a *API) configureRouterField() {
	//users handlers
	a.router.HandleFunc(prefix+"/users/register", a.RegisterUser).Methods("POST")
	a.router.HandleFunc(prefix+"/users/auth", a.Authenticate).Methods("POST")
	a.router.HandleFunc(prefix+"/users/test", a.Authenticate).Methods("GET")

	//admin
	a.router.HandleFunc(prefix+"/admin/createRole", a.CreateRole).Methods("POST")
	a.router.HandleFunc(prefix+"/admin/assignRole", a.AssignRole).Methods("POST")
	a.router.HandleFunc(prefix+"/admin/listRoles", a.ListRoles).Methods("POST")
	a.router.HandleFunc(prefix+"/admin/listPerms/{id}", a.ListRolePerms).Methods("GET")
	a.router.HandleFunc(prefix+"/admin/addPerm", a.addPerm).Methods("POST")
	a.router.HandleFunc(prefix+"/admin/removePerm", a.removePerm).Methods("POST")

	//data handlers
	a.router.HandleFunc(prefix+"/data/{endpoint}/{param}", a.RouteHandler("GET")).Methods("GET")
	a.router.HandleFunc(prefix+"/data/{endpoint}/{param}", a.RouteHandler("POST")).Methods("POST")
	a.router.HandleFunc(prefix+"/data/{endpoint}/{param}", a.RouteHandler("PUT")).Methods("PUT")
	a.router.HandleFunc(prefix+"/data/{endpoint}/{param}", a.RouteHandler("DELETE")).Methods("DELETE")
}
func (a *API) configreStorageField() error {
	storage := storage.New(a.Config.Storage)

	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
