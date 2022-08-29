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
		fmt.Println("accesing data")
		var perm models.Permission
		perm.Path = "/" + mux.Vars(req)["endpoint"]
		if mux.Vars(req)["param"] != "post" {
			perm.Path = perm.Path + "/" + mux.Vars(req)["param"]
		}
		perm.Method = method
		fmt.Println("decoding token")
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
		if err := api.ValidateToken(reqToken, &perm); err != nil {
			api.logger.Info("error validating token: ", err)
			http.Error(writer, "access denied", http.StatusForbidden)
			return
		}
		servers := *api.Config.Servers
		fmt.Println(perm.ConstructUrl(servers[perm.ServerId]))
		fmt.Sprintf("%s %s", perm.Method, perm.ConstructUrl(servers[perm.ServerId]))
		request, err := http.NewRequest(perm.Method, perm.ConstructUrl(servers[perm.ServerId]), req.Body)
		if err != nil {
			fmt.Println("error creating request: ", err)
			http.Error(writer, "internal error", http.StatusInternalServerError)
			return
		}
		fmt.Println("sending request")
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("error sending request: ", err)
			http.Error(writer, "internal error", http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()
		fmt.Println("reading data")
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
