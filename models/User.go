package models

import "golang.org/x/crypto/bcrypt"

//User model defeniton
type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     *Role
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (user *User) GetHashedPassword() (string, error) {
	HashBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return "", err
	}
	return string(HashBytes), nil
}

func (user *User) ValidatePassword(hash []byte) bool {
	if bcrypt.CompareHashAndPassword(hash, []byte(user.Password)) != nil {
		return false
	}
	return true
}
