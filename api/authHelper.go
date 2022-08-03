package api

import (
	"errors"
	"github.com/form3tech-oss/jwt-go"
	"time"
)

type JWTClaim struct {
	UserId int
	jwt.StandardClaims
}

func (api *API) GenerateJWT(userId int) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(api.config.SecretKey))
	return
}

func (api *API) ValidateToken(signedToken string, endpoint string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(api.config.SecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	res, err := api.storage.RolePermRep.CheckPermission(claims.UserId, endpoint)
	if err != nil {
		return err
	}
	if !res {
		err = errors.New("access denied")
	}
	return err
}
