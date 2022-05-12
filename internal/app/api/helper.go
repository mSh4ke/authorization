package api

import (
	"GitHab/Autorization/internal/app/middleware"
	"GitHab/Autorization/storage"
	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
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

	a.router.Handle("/images", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.Getimages))).Methods("GET")
	a.router.Handle("/images", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.Postimage))).Methods("POST")
	a.router.Handle("/images/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.Updateimages))).Methods("PUT")
	a.router.Handle("/images/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.Deleteimages))).Methods("DELETE")

	a.router.Handle("/videos", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.GetVideos))).Methods("GET")
	a.router.Handle("/videos", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.PostVideos))).Methods("POST")
	a.router.Handle("/videos/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.UpdateVideos))).Methods("PUT")
	a.router.Handle("/videos/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.DeleteVideos))).Methods("DELETE")

	a.router.Handle(prefix+"/products", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.Getproducts))).Methods("GET")
	a.router.Handle(prefix+"/products", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.Postproducts))).Methods("POST")
	a.router.Handle(prefix+"/products/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.Updateproducts))).Methods("PUT")
	a.router.Handle(prefix+"/products/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.Deleteproducts))).Methods("DELETE")

	a.router.HandleFunc(prefix+"/brand/get/{id}", a.GetBrandsById).Methods("POST")
	a.router.HandleFunc(prefix+"/brands/get", a.GetBrands).Methods("POST")
	//a.router.Handle(prefix+"/brands", middleware.JwtMiddleware.Handler(
	//	http.HandlerFunc(a.GetBrands))).Methods("GET")
	a.router.Handle(prefix+"/brands", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.PostBrands))).Methods("POST")
	a.router.Handle(prefix+"/brands/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.UpdateBrands))).Methods("PUT")
	a.router.Handle(prefix+"/brands/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.DeleteBrands))).Methods("DELETE")

	a.router.HandleFunc("categories", a.Getcategories).Methods("GET")
	a.router.HandleFunc("categories", a.Postcategory).Methods("POST")
	a.router.HandleFunc("categories/{id}", a.Updatecategory).Methods("PUT")
	a.router.HandleFunc("categories/{id}", a.Deletecategory).Methods("DELETE")
}
func (a *API) configreStorageField() error {
	storage := storage.New(a.config.Storage)

	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
