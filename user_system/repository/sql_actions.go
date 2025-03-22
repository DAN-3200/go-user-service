package repository

import (
	"app/model"
	"database/sql"
	"fmt"
)

// Receber qualquer banco SQL
type SQLManager struct {
	DB *sql.DB
}

func NewSQLManager(db *sql.DB) *SQLManager {
	return &SQLManager{db}
}

// -- SQL Methods

func (it *SQLManager) UserSaveSQL(info model.User) error {
	var query = `INSERT INTO users (name, email, password, role) VALUES ($1,$2,$3,$4)`
	_, err := it.DB.Exec(query, info.Name, info.Email, info.Password, info.Role)
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return err
	}
	return nil
}
func (it *SQLManager) UserReadSQL() ([]model.User, error) {
	return []model.User{
		{Id: 0, Name: "niel", Email: "", Password: "", Role: "", Date: ""},
	}, nil
}
func (it *SQLManager) UserUpdateSQL(info model.User) error {
	return nil
}
func (it *SQLManager) UserDeleteSQL(info struct{ Id string }) error {
	return nil
}

// ----
func (it *SQLManager) CreateUserTable() {
	var _, err = it.DB.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(50) NOT NULL,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
	)
	if err != nil {
		fmt.Println("Erro", err)
		return
	}
}
