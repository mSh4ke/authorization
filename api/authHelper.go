package api

import (
	"errors"
	"github.com/form3tech-oss/jwt-go"
	"github.com/mSh4ke/authorization/models"
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
	tokenString, err = token.SignedString([]byte(api.Config.SecretKey))
	return
}

func (api *API) ValidateToken(signedToken string) (int, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(api.Config.SecretKey), nil
		},
	)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("bad token")
		return 0, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return 0, err
	}
	return claims.UserId, nil
}

func (api *API) ValidatePermission(userId int, perm *models.Permission) error {
	if err := api.storage.RolePermRep.CheckPermission(userId, perm); err != nil {
		return err
	}
	return nil
}
