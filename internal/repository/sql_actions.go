package repository

import (
	"app/internal/dto"
	"app/internal/model"
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

// ------------------------------------------------------------------------

func (it *SQLManager) UserSaveSQL(info model.User) error {
	query := `INSERT INTO users (
		id,
		name, 
		email, 
		password_hash, 
		is_email_verified, 
		is_active,
		created_at,
		updated_at,
		role
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9);`

	_, err := it.DB.Exec(query,
		info.ID,
		info.Name,
		info.Email,
		info.PasswordHash,
		info.IsEmailVerified,
		info.IsActive,
		info.CreatedAt,
		info.UpdatedAt,
		info.Role,
	)
	if err != nil {
		return err
	}
	return nil
}

func (it *SQLManager) UserReadSQL(infoID string) (dto.UserRes, error) {
	query := `SELECT id, name, email, role, created_at FROM users WHERE id=$1`
	row := it.DB.QueryRow(query, infoID)

	var userObj dto.UserRes
	err := row.Scan(
		&userObj.ID,
		&userObj.Name,
		&userObj.Email,
		&userObj.Role,
		&userObj.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Nenhum registro encontro.")
		}
		return dto.UserRes{}, err
	}

	return userObj, nil
}

func (it *SQLManager) ReadAllUserSQL() ([]dto.UserRes, error) {
	query := `SELECT id, name, email, role, created_at FROM users`
	rows, err := it.DB.Query(query)
	if err != nil {
		fmt.Printf("Erro de consulta: %v", err)
		return []dto.UserRes{}, err
	}

	var userList []dto.UserRes
	var m dto.UserRes

	for rows.Next() {
		var err = rows.Scan(
			&m.ID,
			&m.Name,
			&m.Email,
			&m.Role,
			&m.CreatedAt,
		)
		if err != nil {
			fmt.Printf("Erro de Leitura dos dados do Banco: %v", err)
			return []dto.UserRes{}, err
		}
		userList = append(userList, m)
	}
	rows.Close()
	return userList, nil
}

func (it *SQLManager) UserUpdateSQL(info dto.UserUpdateReq) error {
	query := `UPDATE users SET name=$1, password_hash=$2 WHERE id=$3`

	_, err := it.DB.Exec(query, info.Name, info.PasswordHash, info.ID)
	if err != nil {
		return err
	}

	return nil
}

func (it *SQLManager) UserDeleteSQL(infoID string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := it.DB.Exec(query, infoID)
	if err != nil {
		fmt.Println("Erro ao excluir: ", err)
		return err
	}
	return nil
}

func (it *SQLManager) LoginUserSQL(userEmail string) (model.User, error) {
	query := `SELECT id, name, password_hash, email, role FROM users WHERE email=$1;`

	var mU model.User
	var err = it.DB.QueryRow(query, userEmail).
		Scan(
			&mU.ID,
			&mU.Name,
			&mU.PasswordHash,
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

func (it *SQLManager) MyInfoSQL(infoID string) (model.User, error) {
	query := `SELECT id,
		name, 
		email, 
		password_hash, 
		is_email_verified, 
		is_active,
		created_at,
		updated_at,
		role
	FROM users WHERE id=$1`
	row := it.DB.QueryRow(query, infoID)

	var obj model.User
	var err = row.Scan(
		&obj.ID,
		&obj.Name,
		&obj.Email,
		&obj.PasswordHash,
		&obj.IsEmailVerified,
		&obj.IsActive,
		&obj.CreatedAt,
		&obj.UpdatedAt,
		&obj.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Nenhum registro encontro.")
		}
		return model.User{}, err
	}

	return obj, nil
}

func (it *SQLManager) GetUserByEmail(email string) (model.User, error) {
	query := `SELECT id, name, role FROM users WHERE email=$1`
	row := it.DB.QueryRow(query, email)

	var obj model.User
	err := row.Scan(
		&obj.ID,
		&obj.Name,
		&obj.Role,
	)

	if err != nil {
		return model.User{}, err
	}

	return obj, nil
}

func (it *SQLManager) EditMyInfoSQL(info map[string]any) error {
	return nil
}

func (it *SQLManager) RefreshPassword(info dto.RefreshPassword) error {
	query := `UPDATE users SET password_hash=$1 WHERE id=$2`
	_, err := it.DB.Exec(query, info.NewPassword, info.ID)
	if err != nil {
		return err
	}

	return nil
}

func (it *SQLManager) ValidateEmail(email string) error {
	query := `UPDATE users SET is_email_verified=TRUE WHERE email=$1`
	_, err := it.DB.Exec(query, email)
	if err != nil {
		return err
	}

	return nil
}
