package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mSh4ke/authorization/models"
	"log"
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

func (RolePermRep *RolePermRep) AddPermission(tx *sql.Tx, ctx *context.Context, roleId int, permId int) error {
	query := fmt.Sprintf("INSERT INTO %s (roles_id,permissions_id) VALUES ($1,$2)", rolePermTable)
	if _, err := tx.QueryContext(*ctx, query, roleId, permId); err != nil {
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

const ListRolePerms = "SELECT p.id,p.req_path FROM roles_permissions AS rp LEFT JOIN permissions AS p ON rp.permissions_id = p.id WHERE rp.roles_id = $1"

func (RolePermRep *RolePermRep) ListRolePerms(roleId int) (*[]models.Permission, error) {
	rows, err := RolePermRep.storage.db.Query(ListRolePerms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	perms := make([]models.Permission, 0)
	for rows.Next() {
		var perm models.Permission
		if err := rows.Scan(&perm.Id, &perm.Path); err != nil {
			log.Println(err)
			continue
		}
		perms = append(perms, perm)
	}
	return &perms, nil
}

func (RolePermRep *RolePermRep) AssignPermissions(roleId int, permsId *[]int) error {
	ctx := context.Background()
	tx, err := RolePermRep.storage.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	for permId := range *permsId {
		if err := RolePermRep.AddPermission(tx, &ctx, roleId, permId); err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
