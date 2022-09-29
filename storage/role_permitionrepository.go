package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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
		("WHERE u.id = $1 AND p.req_path = $2 AND p.req_method = $3")
	if err := RolePermRep.storage.db.QueryRow(query, userId, perm.ParseUrl(), perm.Method).Scan(&perm.ServerId); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (RolePermRep *RolePermRep) AddPermission(tx *sql.Tx, ctx *context.Context, roleId int, permId int) error {
	query := fmt.Sprintf("INSERT INTO %s (roles_id,permissions_id) VALUES ($1,$2)", rolePermTable)
	if _, err := tx.ExecContext(*ctx, query, roleId, permId); err != nil {
		return err
	}
	return nil
}

const CountRolePerms = "SELECT COUNT(p.id) FROM roles_permissions AS rp LEFT JOIN permissions AS p ON rp.permissions_id = p.id LEFT JOIN roles AS r ON rp.roles_id = r.id"
const ListRolePerms = "SELECT p.id,p.req_path,p.req_method FROM roles_permissions AS rp LEFT JOIN permissions AS p ON rp.permissions_id = p.id LEFT JOIN roles AS r ON rp.roles_id = r.id"

func (RolePermRep *RolePermRep) ListRolePerms(pgReq *models.PageRequest) (*[]models.Permission, error) {
	if err := RolePermRep.storage.db.QueryRow(CountRolePerms + pgReq.PageReq()).Scan(&pgReq.TotalRecords); err != nil {
		return nil, err
	}
	rows, err := RolePermRep.storage.db.Query(ListRolePerms + pgReq.PageReq())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	perms := make([]models.Permission, 0)
	for rows.Next() {
		var perm models.Permission
		if err := rows.Scan(&perm.Id, &perm.Path, &perm.Method); err != nil {
			log.Println(err)
			continue
		}
		perms = append(perms, perm)
	}
	return &perms, nil
}

const ClearRolePerms = "DELETE FROM roles_permissions WHERE roles_id = $1"

func (RolePermRep *RolePermRep) AssignPermissions(roleId int, permsId *[]int) error {
	ctx := context.Background()
	tx, err := RolePermRep.storage.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if _, err = tx.QueryContext(ctx, ClearRolePerms, roleId); err != nil {
		tx.Rollback()
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
