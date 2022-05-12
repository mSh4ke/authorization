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

func (api *API) Getcategories(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer, req)
	var (
		filter models.Filter
	)
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
	request, err := http.NewRequest("GET", "http://localhost:8085/api/v2/categories", bytes.NewBuffer(byteArr))
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
func (api *API) Postcategory(writer http.ResponseWriter, req *http.Request) {
	var category models.Categories

	err := json.NewDecoder(req.Body).Decode(&category)
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
	fmt.Println(&category)
	byteArr, err := json.MarshalIndent(category, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	Bearer := ChekToken()
	request, err := http.NewRequest("POST", "http://localhost:8085/api/v2/categories", bytes.NewBuffer(byteArr))
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
func (api *API) Updatecategory(writer http.ResponseWriter, req *http.Request) {
	var image models.Brand
	err := json.NewDecoder(req.Body).Decode(&image)
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
	fmt.Println(&image)
	byteArr, err := json.MarshalIndent(image, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	Bearer := ChekToken()
	request, err := http.NewRequest("PUT", "http://localhost:8085/api/v2/categories/"+mux.Vars(req)["id"], bytes.NewBuffer(byteArr))
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
func (api *API) Deletecategory(writer http.ResponseWriter, req *http.Request) {
	Bearer := ChekToken()
	request, err := http.NewRequest("DELETE", "http://localhost:8085/api/v2/categories/"+mux.Vars(req)["id"], nil)
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
