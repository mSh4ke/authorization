package storage

import (
	"fmt"
	"github.com/mSh4ke/authorization/models"
)

type roleRepository struct {
	storage *Storage
}

const roleTable string = "roles"

func (roleRep *roleRepository) Create(role *models.Role) error {
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ('$1')", roleTable)
	if _, err := roleRep.storage.db.Query(query, role.Name); err != nil {
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
