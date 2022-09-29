package storage

import (
	"fmt"
	"github.com/mSh4ke/authorization/models"
	"log"
)

type roleRepository struct {
	storage *Storage
}

const roleTable string = "roles"

func (roleRep *roleRepository) Create(role *models.Role) error {
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", roleTable)
	if err := roleRep.storage.db.QueryRow(query, role.Name).Scan(&role.Id); err != nil {
		return err
	}
	return nil
}

func (roleRep *roleRepository) Delete(roleId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d", roleTable, roleId)
	if _, err := roleRep.storage.db.Query(query); err != nil {
		return err
	}
	return nil
}

func (roleRep *roleRepository) Rename(roleId int, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name = %s WHERE id = %d", roleTable, name, roleId)
	if _, err := roleRep.storage.db.Query(query); err != nil {
		return err
	}
	return nil
}

func (roleRep *roleRepository) ListRoles() (*[]models.Role, error) {
	query := fmt.Sprintf("SELECT id,name FROM %s", roleTable)
	Rows, err := roleRep.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer Rows.Close()
	roles := make([]models.Role, 0)
	for Rows.Next() {
		var role models.Role
		if err := Rows.Scan(&role.Id, &role.Name); err != nil {
			log.Println("Error reading role data: ", err)
			continue
		}
		roles = append(roles, role)
	}
	return &roles, nil
}
