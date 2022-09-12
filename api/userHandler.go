package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mSh4ke/authorization/models"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (api *API) Authenticate(wrt http.ResponseWriter, req *http.Request) {
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

func (api *API) ListUsers(wrt http.ResponseWriter, req *http.Request) {
	fmt.Println("accesing usersList")
	perm := models.Permission{
		Path:     "/users/list",
		Method:   "POST",
		ServerId: 0,
	}

	fmt.Println("decoding token")
	reqToken := req.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		log.Println("invalid token")
		log.Println(reqToken)
		http.Error(wrt, "invalid token", http.StatusForbidden)
		return
	}

	reqToken = strings.TrimSpace(splitToken[1])

	if reqToken == "" && api.Config.UnauthorizedId != 0 {
		return
	}
	if err := api.ValidateToken(reqToken, &perm); err != nil {
		api.logger.Info("error validating token: ", err)
		http.Error(wrt, "access denied", http.StatusForbidden)
		return
	}
	fields := make([]models.Field, 0)
	pgReq := models.PageRequest{Fields: &fields}
	err := json.NewDecoder(req.Body).Decode(&pgReq)
	if err != nil {
		http.Error(wrt, "invalid json", 400)
		return
	}
	users, err := api.storage.UserRepository.List(&pgReq)
	if err != nil {
		log.Println("failed fetching user list")
		log.Println(err)
		http.Error(wrt, "internal error", http.StatusInternalServerError)
	}
	json.NewEncoder(wrt).Encode(users)
	wrt.WriteHeader(http.StatusOK)
}

func (api *API) GetUser(wrt http.ResponseWriter, req *http.Request) {
	fmt.Println("accesing user data")
	perm := models.Permission{
		Path:     "users/list",
		Method:   "POST",
		ServerId: 0,
	}
	var uid int
	if id, err := strconv.Atoi(mux.Vars(req)["id"]); err != nil || id == 0 {
		http.Error(wrt, "bad id", http.StatusBadRequest)
	} else {
		uid = id
	}
	fmt.Println("decoding token")
	reqToken := req.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		log.Println("invalid token")
		log.Println(reqToken)
		http.Error(wrt, "invalid token", http.StatusForbidden)
		return
	}

	reqToken = strings.TrimSpace(splitToken[1])

	if reqToken == "" && api.Config.UnauthorizedId != 0 {
		return
	}
	if err := api.ValidateToken(reqToken, &perm); err != nil {
		api.logger.Info("error validating token: ", err)
		http.Error(wrt, "access denied", http.StatusForbidden)
		return
	}
	user, err := api.storage.UserRepository.Get(uid)
	if err != nil {
		log.Println("failed fetching user list")
		log.Println(err)
		http.Error(wrt, "internal error", http.StatusInternalServerError)
	}
	json.NewEncoder(wrt).Encode(user)
	wrt.WriteHeader(http.StatusOK)
}
