package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mSh4ke/authorization/models"
	"log"
	"net/http"
	"strconv"
)

func (api *API) CreateRole(wrt http.ResponseWriter, req *http.Request) {
	perm := models.Permission{
		Path:     "/admin/createRole",
		Method:   "POST",
		ServerId: 0,
	}
	reqToken := req.Header.Get("Authorization")
	fmt.Println(reqToken)
	if err := api.ValidateToken(reqToken, &perm); err != nil {
		api.logger.Info(err)
		http.Error(wrt, "access denied", http.StatusForbidden)
		return
	}
	var role models.Role
	if err := json.NewDecoder(req.Body).Decode(&role); err != nil {
		log.Println("cannot decode body json: ", err)
		http.Error(wrt, "invalid body json", http.StatusBadRequest)
		return
	}

	if err := api.storage.RoleRep.Create(&role); err != nil {
		log.Println("failed creating role: ", err)
		http.Error(wrt, "internal error", http.StatusInternalServerError)
		return
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
	perm := models.Permission{
		Path:     "/admin/assignRole",
		Method:   "Post",
		ServerId: 0,
	}
	reqToken := req.Header.Get("Authorization")
	fmt.Println(reqToken)
	if err := api.ValidateToken(reqToken, &perm); err != nil {
		api.logger.Info(err)
		http.Error(wrt, "access denied", http.StatusForbidden)
		return
	}
	var role models.Role
	if err := json.NewDecoder(req.Body).Decode(&role); err != nil {
		log.Println("cannot decode body json: ", err)
		http.Error(wrt, "invalid body json", http.StatusBadRequest)
		return
	}

	if err := api.storage.RoleRep.Create(&role); err != nil {
		log.Println("failed creating role: ", err)
		http.Error(wrt, "internal error", 500)
		return
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

func (api *API) addPerm(wrt http.ResponseWriter, req *http.Request) {
	perm := models.Permission{
		Path:     "/admin/addPerm",
		Method:   "Post",
		ServerId: 0,
	}
	reqToken := req.Header.Get("Authorization")
	fmt.Println(reqToken)
	if err := api.ValidateToken(reqToken, &perm); err != nil {
		api.logger.Info(err)
		http.Error(wrt, "access denied", 403)
		return
	}
	var rolePerm struct {
		roleId int `json:"role_id"`
		permId int `json:"perm_id"`
	}

	if err := json.NewDecoder(req.Body).Decode(&rolePerm); err != nil {
		log.Println("cannot decode body json: ", err)
		http.Error(wrt, "invalid body json", 400)
		return
	}

	if err := api.storage.RolePermRep.AddPermission(rolePerm.roleId, rolePerm.permId); err != nil {
		log.Println("failed creating role: ", err)
		http.Error(wrt, "internal error", 500)
		return
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

func (api *API) removePerm(wrt http.ResponseWriter, req *http.Request) {
	perm := models.Permission{
		Path:     "/admin/removePerm",
		Method:   "Delete",
		ServerId: 0,
	}
	reqToken := req.Header.Get("Authorization")
	fmt.Println(reqToken)
	if err := api.ValidateToken(reqToken, &perm); err != nil {
		api.logger.Info(err)
		http.Error(wrt, "access denied", 403)
		return
	}
	var rolePerm struct {
		roleId int `json:"role_id"`
		permId int `json:"perm_id"`
	}

	if err := json.NewDecoder(req.Body).Decode(&rolePerm); err != nil {
		log.Println("cannot decode body json: ", err)
		http.Error(wrt, "invalid body json", 400)
		return
	}

	if err := api.storage.RolePermRep.RemovePermission(rolePerm.roleId, rolePerm.permId); err != nil {
		log.Println("failed creating role: ", err)
		http.Error(wrt, "internal error", 500)
		return
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

func (api *API) ListRoles(wrt http.ResponseWriter, req *http.Request) {
	perm := models.Permission{
		Path:     "/admin/listRoles",
		Method:   "Get",
		ServerId: 0,
	}
	reqToken := req.Header.Get("Authorization")
	fmt.Println(reqToken)
	if err := api.ValidateToken(reqToken, &perm); err != nil {
		api.logger.Info(err)
		http.Error(wrt, "access denied", 403)
		return
	}

	roles, err := api.storage.RoleRep.ListRoles()
	if err != nil {
		log.Println("failed creating role: ", err)
		http.Error(wrt, "internal error", 500)
		return
	}
	wrt.WriteHeader(200)
	json.NewEncoder(wrt).Encode(roles)
	return
}

func (api *API) ListRolePerms(wrt http.ResponseWriter, req *http.Request) {
	perm := models.Permission{
		Path:     "/admin/listRolePerms",
		Method:   "Get",
		ServerId: 0,
	}

	reqToken := req.Header.Get("Authorization")
	fmt.Println(reqToken)
	if err := api.ValidateToken(reqToken, &perm); err != nil {
		api.logger.Info(err)
		http.Error(wrt, "access denied", 403)
		return
	}
	roleid, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(wrt, "malformed id", 400)
		return
	}

	roles, err := api.storage.RolePermRep.ListRolePerms(roleid)
	if err != nil {
		log.Println("failed creating role: ", err)
		http.Error(wrt, "internal error", 500)
		return
	}
	wrt.WriteHeader(200)
	json.NewEncoder(wrt).Encode(roles)
	return
}
