package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

const HeaderString = "Bearer"

func (api *API) RouteHandler(method string) func(writer http.ResponseWriter, req *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		initHeaders(writer, req)
		endpoint := mux.Vars(req)["endpoint"]
		reqToken := req.Header.Get(HeaderString)
		fmt.Println(reqToken)

		if err := api.ValidateToken(reqToken, endpoint); err != nil {
			api.logger.Info("error validating token: ", err)
			http.Error(writer, "access denied", 403)
			return
		}

		request, err := http.NewRequest(method, endpoint, req.Body)
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
