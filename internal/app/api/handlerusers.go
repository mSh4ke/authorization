package api

import (
	"GitHab/Autorization/internal/app/middleware"
	"GitHab/Autorization/internal/app/models"
	"encoding/json"
	"github.com/form3tech-oss/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var (
	token     string
	Bearer    string
	Userlogin string
)

type Authorization struct {
	StatusCode int    `json:"status_code"`
	Token      string `json:"message"`
	Login      string `json:"login"`
	Role       int    `json:"role"`
	IsError    bool   `json:"is_error"`
	PS         string `json:"Maxims_mesage""`
}

func (api *API) RegistratedUser(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer, req)
	api.logger.Info("Post Units POST /api/v1/user/registrate")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid json recieved from Units")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = api.storage.Users().RegistrateUsers(&user)
	if err != nil {
		api.logger.Info("Troubles while registrate new user:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again.",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
	}
	msg := Message{
		StatusCode: 201,
		Message:    "User successfully created.",
		IsError:    true,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) PostToAuth(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer, req)
	api.logger.Info("Post to Auth POST /api/v1/user/auth")
	var userFromJson models.User
	err := json.NewDecoder(req.Body).Decode(&userFromJson)

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
	Userlogin = userFromJson.Login
	userInDB, ok, err := api.storage.Users().FindByLogin(userFromJson.Login)
	if err != nil {
		api.logger.Info("Can not make user search in database:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles while accessing database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !ok {
		api.logger.Info("User with that login does not exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login does not exists in database. Try register first",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	chek := CheckPasswordHash(userFromJson.Password, userInDB.Password)

	if chek != true {
		api.logger.Info("Invalid credetials to auth")
		msg := Message{
			StatusCode: 404,
			Message:    "Your password is invalid",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	Token := jwt.New(jwt.SigningMethodHS256)
	claims := Token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix() //Время жизни токена
	claims["admin"] = true
	claims["name"] = userInDB.Login
	TokenString, err := Token.SignedString(middleware.SecretKey)

	//В случае, если токен выбить не удалось!
	if err != nil {
		api.logger.Info("Can not claim jwt-token")
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//В случае, если токен успешно выбит - отдаем его клиенту
	token = TokenString
	msg := Authorization{
		StatusCode: 200,
		Token:      token,
		Login:      userInDB.Login,
		Role:       userInDB.Role,
		IsError:    true,
		PS:         "",
	}

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func ChekToken() string {
	var Value string

	if token == "" {
		Value = ""

	} else {
		Value = "Bearer " + token
	}
	return Value
}
