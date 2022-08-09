package models

import (
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

//User model defeniton
type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     *Role
	Profile  UserProfile
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserProfile struct {
	DisplayName string `json:"display_name"`
	ContactInfo string `json:"contact_info"`
}

func (user *User) ValidateEmail() bool {
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return false
	}
	return true
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
