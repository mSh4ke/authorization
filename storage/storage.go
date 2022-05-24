package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // для того, чотбы отработала функция init()
)

//Instance of storage
type Storage struct {
	config *Config
	db     *sql.DB

	userRepository           *Userrepository
	role_permitionRepository *Role_permitionrepository
}

//Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
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

func (s *Storage) Users() *Userrepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &Userrepository{
		storage: s,
	}
	return s.userRepository
}
func (s *Storage) Role_permitions() *Role_permitionrepository {
	if s.userRepository != nil {
		return s.role_permitionRepository
	}
	s.role_permitionRepository = &Role_permitionrepository{
		storage: s,
	}
	return s.role_permitionRepository
}
