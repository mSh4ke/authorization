package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mSh4ke/authorization/models"
	"io/ioutil"
	"net/http"
)

const HeaderString = "Bearer"

func (api *API) RouteHandler(method string) func(writer http.ResponseWriter, req *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		initHeaders(writer, req)
		var perm models.Permission
		perm.Path = "/" + mux.Vars(req)["endpoint"]
		if mux.Vars(req)["param"] != "" {
			perm.Path = perm.Path + "/" + mux.Vars(req)["param"]
		}
		perm.Method = method
		reqToken := req.Header.Get(HeaderString)
		fmt.Println(reqToken)
		if reqToken == "" && api.Config.UnauthorizedId != 0 {
			return
		}
		if err := api.ValidateToken(reqToken, &perm); err != nil {
			api.logger.Info("error validating token: ", err)
			http.Error(writer, "access denied", 403)
			return
		}
		servers := *api.Config.Servers
		request, err := http.NewRequest(perm.Method, perm.ConstructUrl(servers[perm.ServerId]), req.Body)
		if err != nil {
			fmt.Println("error creating request: ", err)
			http.Error(writer, "internal error", 500)
			return
		}

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("error sending request: ", err)
			http.Error(writer, "internal error", 500)
			return
		}
		defer response.Body.Close()
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("error reading response data: ", err)
			http.Error(writer, "internal error", 500)
			return
		}
		writer.WriteHeader(200)
		writer.Write(responseData)
	}
}
