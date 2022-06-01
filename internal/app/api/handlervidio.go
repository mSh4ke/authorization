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

func (api *API) GetVideos(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	initHeaders(writer, req)
	var (
		filter models.PageRequest
	)
	fl := make([]models.Field, 0)

	filter = models.PageRequest{
		Fields: &fl,
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
	request, err := http.NewRequest("GET", "http://localhost:8085/api/v2/videos", bytes.NewBuffer(byteArr))
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

func (api *API) PostVideos(writer http.ResponseWriter, req *http.Request) {
	var image models.Videos
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
	request, err := http.NewRequest("POST", "http://localhost:8085/api/v2/videos", bytes.NewBuffer(byteArr))
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
func (api *API) UpdateVideos(writer http.ResponseWriter, req *http.Request) {
	var image models.Videos
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
	request, err := http.NewRequest("PUT", "http://localhost:8085/api/v2/videos/"+mux.Vars(req)["id"], bytes.NewBuffer(byteArr))
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
func (api *API) DeleteVideos(writer http.ResponseWriter, req *http.Request) {
	Bearer := ChekToken()
	request, err := http.NewRequest("DELETE", "http://localhost:8085/api/v2/videos/"+mux.Vars(req)["id"], nil)
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
