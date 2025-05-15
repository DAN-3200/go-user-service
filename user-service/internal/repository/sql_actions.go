package repository

import (
	"app/internal/model"
	"app/pkg/security"
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
	hashPassword, err := security.HashPassword(info.Password)
	if err != nil {
		return fmt.Errorf("Error Bycript HashPassword")
	}

	_, err = it.DB.Exec(query, info.Name, info.Email, hashPassword, info.Role)
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return err
	}
	return nil
}

func (it *SQLManager) UserReadSQL(infoID int) (model.User, error) {
	var query = `SELECT id, name, email, password, role, date FROM users WHERE id=$1`
	var row = it.DB.QueryRow(query, infoID)

	var userObj model.User
	var err = row.Scan(
		&userObj.Id,
		&userObj.Name,
		&userObj.Email,
		&userObj.Password,
		&userObj.Role,
		&userObj.Date,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Nenhum registro encontro.")
		} else {
			fmt.Println("Erro de consulta: ", err)
		}
		return model.User{}, err
	}

	return userObj, nil
}

func (it *SQLManager) ReadAllUserSQL() ([]model.User, error) {
	var query = `SELECT id, name, email, password, role, date FROM users`
	var rows, err = it.DB.Query(query)
	if err != nil {
		fmt.Printf("Erro de consulta: %v", err)
		return []model.User{}, err
	}

	var userList []model.User
	var m model.User

	for rows.Next() {
		var err = rows.Scan(
			&m.Id,
			&m.Name,
			&m.Email,
			&m.Password,
			&m.Role,
			&m.Date,
		)
		if err != nil {
			fmt.Printf("Erro de Leitura dos dados do Banco: %v", err)
			return []model.User{}, err
		}
		userList = append(userList, m)
	}
	rows.Close()
	return userList, nil
}

func (it *SQLManager) UserUpdateSQL(info model.User) error {
	var query = `UPDATE users SET name=$1, password=$2 WHERE id = $3`
	newPassword, err := security.HashPassword(info.Password)
	if err != nil {
		return fmt.Errorf("Error Bycript HashPassword")
	}

	_, err = it.DB.Exec(query, info.Name, newPassword, info.Id)
	if err != nil {
		fmt.Println("Erro ao atulizar campo da table: ", err)
		return err
	}

	return nil
}

func (it *SQLManager) UserDeleteSQL(infoID int) error {
	var query = `DELETE FROM users WHERE id = $1`
	var _, err = it.DB.Exec(query, infoID)
	if err != nil {
		fmt.Println("Erro ao excluir: ", err)
		return err
	}
	return nil
}

// criar Login Consulta

func (it *SQLManager) LoginUserSQL(UserEmail string) (model.User, error) {
	var query = `SELECT id, name, password, email, role FROM users WHERE email=$1;`

	var mU model.User
	var err = it.DB.QueryRow(query, UserEmail).
		Scan(
			&mU.Id,
			&mU.Name,
			&mU.Password,
			&mU.Email,
			&mU.Role,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Nenhum registro encontrado")
			return mU, fmt.Errorf("Nenhum registro encontrado")
		} else {
			fmt.Println("Erro de consulta: ", err)
		}
		return mU, err
	}

	return mU, nil
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
