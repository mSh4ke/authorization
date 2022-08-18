package storage

import (
	"errors"
	"fmt"
	"github.com/mSh4ke/authorization/models"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
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
	query := fmt.Sprintf("INSERT INTO %s (login,password,email,role_id,display_name,contact_info) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id", tableUsers)

	if err := userRep.storage.db.QueryRow(query, user.Login, passwordHash, user.Email, user.Role.Id, user.Profile.DisplayName, user.Profile.ContactInfo).Scan(&user.Id); err != nil {
		fmt.Println(query)
		return err
	}
	fmt.Println(query)
	return nil
}

func (userRep *Userrepository) AuthenticateUser(user *models.User) error {
	query := fmt.Sprintf("SELECT u.id, u.password, u.email, u.display_name, u.contact_info, u.role_id, r.name FROM %s AS u ", tableUsers) +
		fmt.Sprintf("LEFT JOIN roles AS r ON u.role_id = r.id ") +
		fmt.Sprintf("WHERE u.login = '%s'", user.Login)
	fmt.Println(query)
	var passwordHash string
	if err := userRep.storage.db.QueryRow(query).Scan(&user.Id, &passwordHash, &user.Email, &user.Profile.DisplayName, &user.Profile.ContactInfo, &user.Role.Id, &user.Role.Name); err != nil {
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

func (userRep *Userrepository) EditProfile(userprf models.UserProfile, userId int) error {
	reqArray := make([]string, 0)
	if userprf.DisplayName != "" {
		reqArray = append(reqArray, "display_name = "+userprf.DisplayName)
	}

	if userprf.ContactInfo != "" {
		reqArray = append(reqArray, "contact_info  = "+userprf.ContactInfo)
	}
	queryString := strings.Join(reqArray, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d", usersTable, queryString, userId)
	if _, err := userRep.storage.db.Query(query); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (userRep *Userrepository) InitAdmin() error {
	query := "SELECT COUNT(*) FROM users"
	var userNum int
	if err := userRep.storage.db.QueryRow(query).Scan(&userNum); err != nil {
		fmt.Println(err)
		return err
	}
	if userNum == 0 {
		fmt.Println("Empty user base, creating admin account")
		adminRole := models.Role{Id: 1}
		admin := models.User{Role: &adminRole}
		fmt.Println("Please input admin login")
		fmt.Scanln(&admin.Login)
		fmt.Println("Please input admin password")
		fmt.Scanln(&admin.Password)
		fmt.Println("Please input admin Email")
		fmt.Scanln(&admin.Email)
		if err := userRep.RegisterUser(&admin); err != nil {
			fmt.Println("Registering admin user failed ", err)
			return err
		}
		if err := userRep.AssignRole(admin.Id, 1); err != nil {
			fmt.Println("Failed assigning admin role ", err)
			return err
		}
		return nil
	}
	return nil
}
