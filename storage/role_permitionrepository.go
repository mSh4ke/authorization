package storage

import (
	"fmt"
)

type Role_permitionrepository struct {
	storage *Storage
}

var usersTable = "Users"

func (brRep *Role_permitionrepository) Role(id_per int, log string) bool {
	query := fmt.Sprintf("Select role_id FROM %s WHERE login = %s ", usersTable, log)
	fmt.Println(query)
	var role string
	var status bool
	rows := brRep.storage.db.QueryRow(query).Scan(&role)

	fmt.Println(rows)
	query = fmt.Sprintf("SELECT CASE WHEN EXISTS (Select * FROM roles_permisions WHERE roles_id = %s and permisions_id = %s) THEN 'TRUE' ELSE 'FALSE' END", role, id_per)
	fmt.Println(query)
	rows = brRep.storage.db.QueryRow(query).Scan(&status)
	return status
}
