package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // для того, чотбы отработала функция init()
)

//Instance of storage
type Storage struct {
	config         *Config
	db             *sql.DB
	UserRepository *Userrepository
	RolePermRep    *RolePermRep
	RoleRep        *roleRepository
}

//Storage Constructor
func New(config *Config) *Storage {
	strg := &Storage{config: config}

	strg.UserRepository = &Userrepository{storage: strg}
	strg.RolePermRep = &RolePermRep{storage: strg}
	strg.RoleRep = &roleRepository{storage: strg}

	return strg
}

//Open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection created successfully!")
	return nil
}

//Close connection
func (storage *Storage) Close() {
	storage.db.Close()
}
