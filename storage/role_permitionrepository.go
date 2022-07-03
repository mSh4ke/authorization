package storage

import (
	"fmt"
)

type Role_permitionrepository struct {
	storage *Storage
}

var usersTable = "Users"

func (roleRep *Role_permitionrepository) Role(id_per int, log string) bool {
	var role int
	var status bool
	query := fmt.Sprintf("Select role_id FROM %s WHERE login = '$1' ", usersTable)
	fmt.Println(query)

	if err := roleRep.storage.db.QueryRow(query, log).Scan(&role); err != nil {
		fmt.Println(err)
	}

	query = fmt.Sprintf("SELECT CASE WHEN EXISTS (Select * FROM roles_permisions WHERE roles_id = %d and permisions_id = %d) THEN 'TRUE' ELSE 'FALSE' END", role, id_per)
	fmt.Println(query)
	if err := roleRep.storage.db.QueryRow(query).Scan(&status); err != nil {
		fmt.Println(err)
	}
	return status
}
