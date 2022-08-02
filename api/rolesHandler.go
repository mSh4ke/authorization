package api

import (
	"encoding/json"
	"fmt"
	"github.com/mSh4ke/authorization/models"
	"log"
	"net/http"
)

func (api *API) CreateRole(wrt http.ResponseWriter, req *http.Request) {
	initHeaders(wrt, req)
	PermString := "/admin/createRole"
	reqToken := req.Header.Get(HeaderString)
	fmt.Println(reqToken)
	if err := api.ValidateToken(reqToken, PermString); err != nil {
		api.logger.Info(err)
		http.Error(wrt, "access denied", 403)
		return
	}
	var role models.Role
	if err := json.NewDecoder(req.Body).Decode(&role); err != nil {
		log.Println("cannot decode body json: ", err)
		http.Error(wrt, "invalid body json", 400)
	}

	if err := api.storage.RoleRep.Create(&role); err != nil {
		log.Println("failed creating role: ", err)
		http.Error(wrt, "internal error", 500)
	}
	msg := Message{
		StatusCode: 200,
		Message:    "success",
		IsError:    false,
	}
	wrt.WriteHeader(200)
	json.NewEncoder(wrt).Encode(msg)
	return
}

func (api *API) AssignRole(wrt http.ResponseWriter, req *http.Request) {
	initHeaders(wrt, req)
	PermString := "/admin/assignRole"
	reqToken := req.Header.Get(HeaderString)
	fmt.Println(reqToken)
	if err := api.ValidateToken(reqToken, PermString); err != nil {
		api.logger.Info(err)
		http.Error(wrt, "access denied", 403)
		return
	}
	var role models.Role
	if err := json.NewDecoder(req.Body).Decode(&role); err != nil {
		log.Println("cannot decode body json: ", err)
		http.Error(wrt, "invalid body json", 400)
	}

	if err := api.storage.RoleRep.Create(&role); err != nil {
		log.Println("failed creating role: ", err)
		http.Error(wrt, "internal error", 500)
	}
	msg := Message{
		StatusCode: 200,
		Message:    "success",
		IsError:    false,
	}
	wrt.WriteHeader(200)
	json.NewEncoder(wrt).Encode(msg)
	return
}
