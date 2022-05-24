package api

import (
	"GitHab/Autorization/internal/app/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func (api *API) Getproducts(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer, req)
	var (
		filter models.Filter
	)
	fmt.Println("Start operation products")
	pg := models.Pages{}
	fl := make([]models.FieldFilter, 0)
	so := make([]models.FieldSort, 0)
	filter = models.Filter{
		Fields: &fl,
		Sorts:  &so,
		Pages:  &pg,
	}
	err := json.NewDecoder(req.Body).Decode(&filter)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	byteArr, err := json.MarshalIndent(filter, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	Bearer := ChekToken()
	request, err := http.NewRequest("GET", "http://localhost:8085/api/v2/products", bytes.NewBuffer(byteArr))
	if err != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	fmt.Println(request)
	request.Header.Set("Authorization", Bearer)
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "Error write respons",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	defer response.Body.Close()
	fmt.Println(response)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "Error reading response",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.Write(body)
}
func (api *API) GetProductById(writer http.ResponseWriter, req *http.Request) {
	request, err := http.NewRequest("GET", "http://localhost:8085/api/v2/product/"+mux.Vars(req)["id"], nil)
	if err != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "Error write respons",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	defer response.Body.Close()
	fmt.Println(response)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "Error reading response",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.Write(body)
}
func (api *API) Postproducts(writer http.ResponseWriter, req *http.Request) {
	var brand models.Brand

	prod := models.Product{
		Brand: &brand,
	}

	err := json.NewDecoder(req.Body).Decode(&prod)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	fmt.Println(&prod)
	byteArr, err := json.MarshalIndent(prod, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	request, err := http.NewRequest("POST", "http://localhost:8085/api/v2/products", bytes.NewBuffer(byteArr))
	if err != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "Error write respons",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	defer response.Body.Close()
	fmt.Println(response)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "Error reading response",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.Write(body)
}
func (api *API) Updateproducts(writer http.ResponseWriter, req *http.Request) {
	var brand models.Brand
	prod := models.Product{
		Brand: &brand,
	}
	err := json.NewDecoder(req.Body).Decode(&prod)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	fmt.Println(&prod)
	byteArr, err := json.MarshalIndent(prod, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	Bearer := ChekToken()
	request, err := http.NewRequest("PUT", "http://localhost:8085/api/v2/products/"+mux.Vars(req)["id"], bytes.NewBuffer(byteArr))
	if err != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	request.Header.Set("Authorization", Bearer)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "Error write respons",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	defer response.Body.Close()
	fmt.Println(response)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "Error reading response",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.Write(body)
}
func (api *API) Deleteproducts(writer http.ResponseWriter, req *http.Request) {
	Bearer := ChekToken()
	request, err := http.NewRequest("DELETE", "http://localhost:8085/api/v2/products/"+mux.Vars(req)["id"], nil)
	if err != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	request.Header.Set("Authorization", Bearer)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "Error write respons",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	defer response.Body.Close()
	fmt.Println(response)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		msg := Message{
			StatusCode: 500,
			Message:    "Error reading response",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.Write(body)
}
