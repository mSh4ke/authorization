package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mSh4ke/authorization/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (api *API) RouteHandler(method string) func(writer http.ResponseWriter, req *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		var perm models.Permission

		perm.Path = "/" + mux.Vars(req)["endpoint"]
		if mux.Vars(req)["param"] != "post" {
			perm.Path = perm.Path + "/" + mux.Vars(req)["param"]
		}
		perm.Method = method
		reqToken := req.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			log.Println("invalid token")
			log.Println(reqToken)
			http.Error(writer, "invalid token", http.StatusForbidden)
			return
		}

		reqToken = strings.TrimSpace(splitToken[1])

		if reqToken == "" && api.Config.UnauthorizedId != 0 {
			return
		}
		userId, err := api.ValidateToken(reqToken)
		if err != nil {
			api.logger.Info("error validating token: ", err)
			http.Error(writer, err.Error(), http.StatusForbidden)
			return
		}
		err = api.ValidatePermission(userId, &perm)
		if err != nil {
			api.logger.Info("error getting permission", err)
			http.Error(writer, "access denied", http.StatusForbidden)
			return
		}
		servers := *api.Config.Servers
		api.logger.Info("accessing", perm.ConstructUrl(servers[perm.ServerId]))
		fmt.Sprintf("%s %s", perm.Method, perm.ConstructUrl(servers[perm.ServerId]))
		request, err := http.NewRequest(perm.Method, perm.ConstructUrl(servers[perm.ServerId]), req.Body)
		if err != nil {
			api.logger.Info("error creating request: ", err)
			http.Error(writer, "internal error", http.StatusInternalServerError)
			return
		}
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("error sending request: ", err)
			http.Error(writer, "internal error", http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("error reading response data: ", err)
			http.Error(writer, "internal error", http.StatusInternalServerError)
			return
		}
		fmt.Println("writing response")
		writer.WriteHeader(http.StatusOK)
		writer.Write(responseData)
	}
}
