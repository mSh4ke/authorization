package middleware

import (
	"GitHab/Autorization/storage"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"net/http"
)

var Userlogin string
var (
	SecretKey      []byte      = []byte("UltraApi")
	emptyValidFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	}
)

type errorPermition func(w http.ResponseWriter, r *http.Request, err string)

var ErrorPermition errorPermition

func OnError(w http.ResponseWriter, r *http.Request, err string) {
	http.Error(w, "Access denied", http.StatusUnauthorized)
}

func Middleware(id_permition int) *jwtmiddleware.JWTMiddleware {

	var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: emptyValidFunc,
		SigningMethod:       jwt.SigningMethodHS256,
	})
	stat := storage.Role_permitionrepository{}

	if Userlogin == "" {
		Userlogin = "Admin102"
	}
	sta := stat.Role(id_permition, Userlogin)
	if sta == false {
		ErrorPermition = OnError
	}

	return JwtMiddleware
}
