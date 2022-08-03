package api

import (
	"encoding/json"
	"github.com/mSh4ke/authorization/models"
	"log"
	"net/http"
)

func (api *API) Authenticate(wrt http.ResponseWriter, req *http.Request) {
	initHeaders(wrt, req)
	role := models.Role{}
	user := models.User{
		Role: &role,
	}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(wrt, "invalid json", 400)
		return
	}
	log.Println("user data decoded")
	log.Println(user)

	err = api.storage.UserRepository.AuthenticateUser(&user)
	if err != nil {
		log.Println("authentication failed")
		log.Println(err)
		http.Error(wrt, "password is invalid or user does not exist", 400)
		return
	}

	tokenString, err := api.GenerateJWT(user.Id)
	if err != nil {
		log.Println("failed generating token")
		log.Println(err)
		http.Error(wrt, "internal error", 500)
		return
	}

	resp := struct {
		Token string
		User  *models.User
	}{
		Token: tokenString,
		User:  &user,
	}

	json.NewEncoder(wrt).Encode(&resp)
	wrt.WriteHeader(200)
}

func (api *API) RegisterUser(wrt http.ResponseWriter, req *http.Request) {
	initHeaders(wrt, req)

	role := models.Role{}
	user := models.User{
		Role: &role,
	}

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(wrt, "invalid json", 400)
		return
	}
	log.Println("user decoded")
	log.Println("Default role: ", api.Config.DefaultRoleId)
	user.Role.Id = api.Config.DefaultRoleId
	log.Println(user)
	log.Println(role)
	if err != nil {
		log.Println("failed calculating password hash")
		log.Println(err)
		http.Error(wrt, "internal error", 500)
		return
	}
	err = api.storage.UserRepository.RegisterUser(&user)
	if err != nil {
		log.Println("user creation failed")
		log.Println(err)
		http.Error(wrt, "invalid user data", 400)
		return
	}
	wrt.WriteHeader(201)
	json.NewEncoder(wrt).Encode(&user)
}
