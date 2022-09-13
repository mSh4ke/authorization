package storage

import (
	"fmt"
	"github.com/mSh4ke/authorization/models"
)

type RolePermRep struct {
	storage *Storage
}

const usersTable = "users"
const rolePermTable = "roles_permissions"
const permTable = "permissions"

func (RolePermRep *RolePermRep) CheckPermission(userId int, perm *models.Permission) error {
	query := fmt.Sprintf("SELECT p.req_server_id FROM %s AS rp ", rolePermTable) +
		fmt.Sprintf("INNER JOIN %s AS u on u.role_id = rp.roles_id ", usersTable) +
		fmt.Sprintf("INNER JOIN %s AS p on p.id = rp.permissions_id ", permTable) +
		fmt.Sprintf("WHERE u.id = $1 AND p.req_path = $2 AND p.req_method = $3")
	if err := RolePermRep.storage.db.QueryRow(query, userId, perm.ParseUrl(), perm.Method).Scan(&perm.ServerId); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (RolePermRep *RolePermRep) AddPermission(roleId int, permId int) error {
	query := fmt.Sprintf("INSERT INTO %s (roles_id,permissions_id) VALUES ($1,$2)", rolePermTable)
	if _, err := RolePermRep.storage.db.Query(query, roleId, permId); err != nil {
		return err
	}
	return nil
}

func (RolePermRep *RolePermRep) RemovePermission(roleId int, permId int) error {
	query := fmt.Sprintf("DELETE FROM %s (roles_id,permissions_id) VALUES ($1,$2)", rolePermTable)
	if _, err := RolePermRep.storage.db.Query(query, roleId, permId); err != nil {
		return err
	}
	return nil
}
