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

	/*a.router.Handle("/images", middleware.Middleware(13).Handler(
		http.HandlerFunc(a.Getimages))).Methods("GET")
	a.router.Handle("/images", middleware.Middleware(14).Handler(
		http.HandlerFunc(a.Postimage))).Methods("POST")
	a.router.Handle("/images/{id}", middleware.Middleware(15).Handler(
		http.HandlerFunc(a.Updateimages))).Methods("PUT")
	a.router.Handle("/images/{id}", middleware.Middleware(16).Handler(
		http.HandlerFunc(a.Deleteimages))).Methods("DELETE")



		a.router.Handle("/videos", middleware.Middleware().Handler(
			http.HandlerFunc(a.GetVideos))).Methods("GET")
		a.router.Handle("/videos", middleware.Middleware().Handler(
			http.HandlerFunc(a.PostVideos))).Methods("POST")
		a.router.Handle("/videos/{id}", middleware.Middleware().Handler(
			http.HandlerFunc(a.UpdateVideos))).Methods("PUT")
		a.router.Handle("/videos/{id}", middleware.Middleware().Handler(
			http.HandlerFunc(a.DeleteVideos))).Methods("DELETE")

		a.router.HandleFunc(prefix+"/product/get/{id}", a.GetProductById).Methods("POST")
		a.router.Handle(prefix+"/products/get", middleware.Middleware().Handler(
			http.HandlerFunc(a.Getproducts))).Methods("POST")
		a.router.Handle(prefix+"/products/create", middleware.Middleware().Handler(
			http.HandlerFunc(a.Postproducts))).Methods("POST")
		a.router.Handle(prefix+"/product/update/{id}", middleware.Middleware().Handler(
			http.HandlerFunc(a.Updateproducts))).Methods("POST")
		a.router.Handle(prefix+"/product/delete/{id}", middleware.Middleware().Handler(
			http.HandlerFunc(a.Deleteproducts))).Methods("POST")
	*/
	a.router.HandleFunc(prefix+"/brand/get/{id}", a.GetBrandsById).Methods("POST")
	a.router.Handle(prefix+"/brands/get", middleware.Middleware(13).Handler(
		http.HandlerFunc(a.GetBrands))).Methods("POST")
	a.router.Handle(prefix+"/brands/post", middleware.Middleware(14).Handler(
		http.HandlerFunc(a.PostBrands))).Methods("POST")
	a.router.Handle(prefix+"/brands/update/{id}", middleware.Middleware(15).Handler(
		http.HandlerFunc(a.UpdateBrands))).Methods("POST")
	a.router.Handle(prefix+"/brand/delete/{id}", middleware.Middleware(16).Handler(
		http.HandlerFunc(a.DeleteBrands))).Methods("POST")
	/*
		a.router.HandleFunc(prefix+"/categories/get/{id}", a.GetCategoryById).Methods("POST")
		a.router.Handle(prefix+"/categories/get", middleware.Middleware().Handler(
			http.HandlerFunc(a.Getcategories))).Methods("POST")
		a.router.Handle(prefix+"/categories/post", middleware.Middleware().Handler(
			http.HandlerFunc(a.Postcategory))).Methods("POST")
		a.router.Handle(prefix+"/categories/update/{id}", middleware.Middleware().Handler(
			http.HandlerFunc(a.Updatecategory))).Methods("POST")
		a.router.Handle(prefix+"/categories/delete/{id}", middleware.Middleware().Handler(
			http.HandlerFunc(a.Deletecategory))).Methods("POST")

		a.router.Handle(prefix+"/units/get", middleware.Middleware().Handler(
			http.HandlerFunc(a.GetBrands))).Methods("POST")
		a.router.Handle(prefix+"/units/post", middleware.Middleware().Handler(
			http.HandlerFunc(a.PostBrands))).Methods("POST")
		a.router.Handle(prefix+"/units/update/{id}", middleware.Middleware().Handler(
			http.HandlerFunc(a.UpdateBrands))).Methods("POST")
		a.router.Handle(prefix+"/units/delete/{id}", middleware.Middleware().Handler(
			http.HandlerFunc(a.DeleteBrands))).Methods("POST")

		a.router.HandleFunc(prefix+"/companies/get/{id}", a.GetCompaniesById).Methods("POST")
		a.router.Handle(prefix+"/companies/get", middleware.Middleware().Handler(
			http.HandlerFunc(a.GetCompanies))).Methods("POST")
		a.router.Handle(prefix+"/companies/post", middleware.Middleware().Handler(
			http.HandlerFunc(a.PostCompanies))).Methods("POST")
		a.router.Handle(prefix+"/companies/update/{id}", middleware.Middleware().Handler(
			http.HandlerFunc(a.UpdateCompanies))).Methods("POST")
		a.router.Handle(prefix+"/companies/delete/{id}", middleware.Middleware().Handler(
			http.HandlerFunc(a.DeleteCompanies))).Methods("POST")
	*/
}
func (a *API) configreStorageField() error {
	storage := storage.New(a.config.Storage)

	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
