package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/mSh4ke/authorization/models"
)

func (api *API) CreateRole(wrt http.ResponseWriter, req *http.Request) {
	perm := models.Permission{
		Path:     "/admin/createRole",
		Method:   "POST",
		ServerId: 0,
	}
	reqToken := req.Header.Get("Authorization")
	fmt.Println(reqToken)
	userId, err := api.ValidateToken(reqToken)
	if err != nil {
		api.logger.Info(err)
		http.Error(wrt, err.Error(), http.StatusForbidden)
		return
	}
	err = api.ValidatePermission(userId, &perm)
	if err != nil {
		api.logger.Info(err)
		http.Error(wrt, "access denied", http.StatusForbidden)
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
	wrt.WriteHeader(http.StatusOK)
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
	userId, err := api.ValidateToken(reqToken)
	if err != nil {
		api.logger.Info(err)
		http.Error(wrt, err.Error(), http.StatusForbidden)
		return
	}
	err = api.ValidatePermission(userId, &perm)
	if err != nil {
		api.logger.Info(err)
		http.Error(wrt, "access denied", http.StatusForbidden)
	}
	var userRole struct {
		RoleId int `json:"role_id"`
		UserId int `json:"user_id"`
	}
	if err := json.NewDecoder(req.Body).Decode(&userRole); err != nil {
		log.Println("cannot decode body json: ", err)
		http.Error(wrt, "invalid body json", http.StatusBadRequest)
		return
	}

	if err := api.storage.UserRepository.AssignRole(userRole.UserId, userRole.RoleId); err != nil {
		log.Println("failed creating role: ", err)
		http.Error(wrt, "internal error", http.StatusInternalServerError)
		return
	}
	msg := Message{
		StatusCode: 200,
		Message:    "success",
		IsError:    false,
	}
	wrt.WriteHeader(http.StatusOK)
	json.NewEncoder(wrt).Encode(msg)
	return
}

func (api *API) AssignPerm(wrt http.ResponseWriter, req *http.Request) {
	perm := models.Permission{
		Path:     "/admin/addPerm",
		Method:   "Post",
		ServerId: 0,
	}
	reqToken := req.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		log.Println("invalid token")
		log.Println(reqToken)
		http.Error(wrt, "invalid token", http.StatusForbidden)
		return
	}

	reqToken = strings.TrimSpace(splitToken[1])
	fmt.Println(reqToken)
	userId, err := api.ValidateToken(reqToken)
	if err != nil {
		api.logger.Info(err)
		http.Error(wrt, err.Error(), http.StatusForbidden)
		return
	}
	err = api.ValidatePermission(userId, &perm)
	if err != nil {
		api.logger.Info(err)
		http.Error(wrt, "access denied", http.StatusForbidden)
	}
	var rolePerm struct {
		RoleId  int   `json:"role_id"`
		PermIds []int `json:"perm_id"`
	}
	rolePerm.PermIds = make([]int, 0)

	if err := json.NewDecoder(req.Body).Decode(&rolePerm); err != nil {
		log.Println("cannot decode body json: ", err)
		http.Error(wrt, "invalid body json", http.StatusBadRequest)
		return
	}

	if err := api.storage.RolePermRep.AssignPermissions(rolePerm.RoleId, &rolePerm.PermIds); err != nil {
		log.Println("failed adding permissions: ", err)
		http.Error(wrt, "internal error", http.StatusInternalServerError)
		return
	}
	msg := Message{
		StatusCode: 200,
		Message:    "success",
		IsError:    false,
	}
	wrt.WriteHeader(http.StatusOK)
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
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		log.Println("invalid token")
		log.Println(reqToken)
		http.Error(wrt, "invalid token", http.StatusForbidden)
		return
	}

	reqToken = strings.TrimSpace(splitToken[1])
	userId, err := api.ValidateToken(reqToken)
	if err != nil {
		api.logger.Info("error validating token: ", err)
		http.Error(wrt, err.Error(), http.StatusForbidden)
		return
	}
	err = api.ValidatePermission(userId, &perm)
	if err != nil {
		api.logger.Info("error getting permission", err)
		http.Error(wrt, "access denied", http.StatusForbidden)
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
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		log.Println("invalid token")
		log.Println(reqToken)
		http.Error(wrt, "invalid token", http.StatusForbidden)
		return
	}

	reqToken = strings.TrimSpace(splitToken[1])
	fmt.Println(reqToken)
	userId, err := api.ValidateToken(reqToken)
	if err != nil {
		api.logger.Info(err)
		http.Error(wrt, err.Error(), http.StatusForbidden)
		return
	}
	err = api.ValidatePermission(userId, &perm)
	if err != nil {
		api.logger.Info(err)
		http.Error(wrt, "access denied", http.StatusForbidden)
	}
	pgReq := models.PageRequest{}.New()
	err = json.NewDecoder(req.Body).Decode(pgReq)
	if err != nil {
		http.Error(wrt, "malformed id", http.StatusBadRequest)
		return
	}

	roles, err := api.storage.RolePermRep.ListRolePerms(pgReq)
	if err != nil {
		log.Println("failed creating role: ", err)
		http.Error(wrt, "internal error", http.StatusInternalServerError)
		return
	}
	wrt.WriteHeader(http.StatusOK)
	json.NewEncoder(wrt).Encode(roles)
	return
}
