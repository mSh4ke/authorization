package storage

import (
	"fmt"
)

type RolePermRep struct {
	storage *Storage
}

const usersTable = "users"
const rolePermTable = "roles_permisions"
const permTable = "permisions"

func (rolePermRep *RolePermRep) CheckPermission(userId int, permString string) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s AS rp ", rolePermTable) +
		fmt.Sprintf("INNER JOIN %s AS u on u.role_id = rp.roles_id ", usersTable) +
		fmt.Sprintf("INNER JOIN %s AS p on p.name = rp.permisions_id ", permTable) +
		fmt.Sprintf("WHERE u.id = $1 and permString = $2)")
	var result bool
	if err := rolePermRep.storage.db.QueryRow(query, userId, permString).Scan(&result); err != nil {
		fmt.Println(err)
		return false, err
	}
	return result, nil
}

func (RolePermRep *RolePermRep) AddPermission(roleId int, permId int) error {
	query := fmt.Sprintf("INSERT INTO %s (roles_id,permisions_id) VALUES ($1,$2)", rolePermTable)
	if _, err := RolePermRep.storage.db.Query(query, roleId, permId); err != nil {
		return err
	}
	return nil
}

func (RolePermRep *RolePermRep) RemovePermission(roleId int, permId int) error {
	query := fmt.Sprintf("DELETE FROM %s (roles_id,permisions_id) VALUES ($1,$2)", rolePermTable)
	if _, err := RolePermRep.storage.db.Query(query, roleId, permId); err != nil {
		return err
	}
	return nil
}
