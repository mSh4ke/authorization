package storage

import (
	"GitHab/Autorization/internal/app/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Userrepository struct {
	storage *Storage
}

var (
	tableUsers string = "Users"
)

func (userRep *Userrepository) RegistrateUsers(user *models.User) (*models.User, error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	query := fmt.Sprintf("INSERT INTO %s (login,password,role_id) VALUES ($1,$2,$3) RETURNING id", tableUsers)

	if err := userRep.storage.db.QueryRow(query, user.Login, bytes, user.Role).Scan(&user.Id); err != nil {
		fmt.Println(query)
		return nil, err
	}
	fmt.Println(query)
	return user, nil
}

func (ur *Userrepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var userFinded *models.User
	for _, u := range users {
		if u.Login == login {
			userFinded = u
			founded = true
			break
		}
	}
	return userFinded, founded, nil
}

func (ur *Userrepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUsers)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.Id, &u.Login, &u.Password, &u.Role)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil

}
