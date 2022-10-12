package api

import (
	_ "github.com/gorilla/mux"
	"github.com/mSh4ke/authorization/storage"
	"github.com/sirupsen/logrus"
)

var (
	prefix = "/api/v2"
)

func (api *API) configureLoggerField() error {

	log_level, err := logrus.ParseLevel(api.Config.LoggerLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(log_level)
	return nil
}
func (api *API) configureRouterField() {
	//users handlers
	api.router.HandleFunc(prefix+"/users/register", api.RegisterUser).Methods("POST")
	api.router.HandleFunc(prefix+"/users/auth", api.Authenticate).Methods("POST")
	api.router.HandleFunc(prefix+"/users/list", api.ListUsers).Methods("POST")
	api.router.HandleFunc(prefix+"/users/{id}", api.GetUser).Methods("GET")

	//admin
	api.router.HandleFunc(prefix+"/admin/createRole", api.CreateRole).Methods("POST")
	api.router.HandleFunc(prefix+"/admin/assignRole", api.AssignRole).Methods("POST")
	api.router.HandleFunc(prefix+"/admin/listRoles", api.ListRoles).Methods("POST")
	api.router.HandleFunc(prefix+"/admin/listPerms", api.ListPerms).Methods("POST")
	api.router.HandleFunc(prefix+"/admin/assignPerm", api.AssignPerm).Methods("POST")

	//data handlers
	api.router.HandleFunc(prefix+"/data/documents/{endpoint}/{param}", api.RouteHandler("PUT", "documents")).Methods("PUT") //костыль 1
	api.router.HandleFunc(prefix+"/data/documents/{endpoint}/{param}", api.RouteHandler("PUT", "documents")).Methods("PUT") //костыль 2
	api.router.HandleFunc(prefix+"/data/{endpoint}/{param}", api.RouteHandler("GET", "")).Methods("GET")
	api.router.HandleFunc(prefix+"/data/{endpoint}/{param}", api.RouteHandler("POST", "")).Methods("POST")
	api.router.HandleFunc(prefix+"/data/{endpoint}/{param}", api.RouteHandler("PUT", "")).Methods("PUT")
	api.router.HandleFunc(prefix+"/data/{endpoint}/{param}", api.RouteHandler("DELETE", "")).Methods("DELETE")

}
func (api *API) configureStorageField() error {
	storage := storage.New(api.Config.Storage)

	if err := storage.Open(); err != nil {
		return err
	}
	api.storage = storage
	return nil
}
