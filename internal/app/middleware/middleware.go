package middleware

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

var (
	SecretKey      []byte      = []byte("UltraApi")
	emptyValidFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	}
)

func Middleware() *jwtmiddleware.JWTMiddleware {

	var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: emptyValidFunc,
		SigningMethod:       jwt.SigningMethodHS256,
	})
	return JwtMiddleware

}
