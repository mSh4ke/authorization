package storage

import (
	"errors"
	"fmt"
	"github.com/mSh4ke/authorization/models"
	"github.com/sirupsen/logrus"
)

type Userrepository struct {
	storage *Storage
}

const tableUsers string = "Users"

func (userRep *Userrepository) RegisterUser(user *models.User) error {
	passwordHash, err := user.GetHashedPassword()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (login,password,role_id) VALUES ($1,$2,$1)", tableUsers)

	if _, err := userRep.storage.db.Query(query, user.Login, passwordHash, user.Role.Id); err != nil {
		fmt.Println(query)
		return err
	}
	fmt.Println(query)
	return nil
}

func (userRep *Userrepository) AuthenticateUser(user *models.User) error {
	query := fmt.Sprintf("SELECT u.id, u.password, u.role_id, r.name FROM %s AS u ", tableUsers) +
		fmt.Sprintf("LEFT JOIN roles AS r ON u.role_id = r.id ") +
		fmt.Sprintf("WHERE u.login = '%s'", user.Login)
	fmt.Println(query)
	var passwordHash string
	if err := userRep.storage.db.QueryRow(query).Scan(&user.Id, &passwordHash, &user.Role.Id, &user.Role.Name); err != nil {
		logrus.Info(err)
		return err
	}
	if !user.ValidatePassword([]byte(passwordHash)) {
		return errors.New("Invalid password")
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
