package storage

import (
	"fmt"
	"github.com/mSh4ke/authorization/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Userrepository struct {
	storage *Storage
}

const tableUsers string = "Users"

func (userRep *Userrepository) RegisterUser(user *models.User) error {
	query := fmt.Sprintf("INSERT INTO %s (login,password,role_id) VALUES ($1,$2,1)", tableUsers)

	if _, err := userRep.storage.db.Query(query, user.Login, user.Password); err != nil {
		fmt.Println(query)
		return err
	}
	fmt.Println(query)
	return nil
}

func (userRep *Userrepository) AuthenticateUser(user *models.User) error {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	query := fmt.Sprintf("SELECT u.id, u.role_id, r.name FROM %s AS u WHERE login = %s AND password = %s", tableUsers, user.Login, bytes) +
		fmt.Sprintf("LEFT JOIN roles AS r ON u.role_id = r.id")
	fmt.Println(query)
	if err := userRep.storage.db.QueryRow(query).Scan(&user.Id, &user.Role.Id, &user.Role.Name); err != nil {
		logrus.Info(err)
		return err
	}
	return nil
}

func (userRep *Userrepository) AssignRole(userId int, roleId int) error {
	query := fmt.Sprintf("UPDATE %s SET role_id  = %d WHERE id = %d", tableUsers, roleId, userId)
	fmt.Println(query)
	if _, err := userRep.storage.db.Query(query); err != nil {
		logrus.Info(err)
		return err
	}
	return nil
}
