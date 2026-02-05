package repository

import (
	"app/internal/application/dto"
	"app/internal/domain/entity"
	"app/internal/infrastructure/persistence/schema"
	"database/sql"
	"fmt"
)

// Receber qualquer banco SQL
type UserAuthRepo struct {
	DB *sql.DB
}

func NewUserAuthRepo(db *sql.DB) *UserAuthRepo {
	return &UserAuthRepo{db}
}

// ------------------------------------------------------------------------

func (it *UserAuthRepo) LoginUser(userEmail string) (*entity.User, error) {
	query := `SELECT id, name, password_hash, email, role FROM users WHERE email=$1;`

	var mU schema.User
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
			return nil, fmt.Errorf("Nenhum registro encontrado")
		} else {
			fmt.Println("Erro de consulta: ", err)
		}
		return nil, err
	}

	res, err := schema.ReconstituteUser(&mU)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (it *UserAuthRepo) GetUserByEmail(email string) (*entity.User, error) {
	query := `SELECT id, name, role FROM users WHERE email=$1`
	row := it.DB.QueryRow(query, email)

	var obj schema.User
	err := row.Scan(
		&obj.ID,
		&obj.Name,
		&obj.Role,
	)

	if err != nil {
		return nil, err
	}
	
	res, err := schema.ReconstituteUser(&obj)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (it *UserAuthRepo) RefreshPassword(id string, info dto.RefreshPassword) error {
	query := `UPDATE users SET password_hash=$1 WHERE id=$2`
	_, err := it.DB.Exec(query, info.NewPassword, id)
	if err != nil {
		return err
	}

	return nil
}

func (it *UserAuthRepo) ValidateEmail(email string) error {
	query := `UPDATE users SET is_email_verified=TRUE WHERE email=$1`
	_, err := it.DB.Exec(query, email)
	if err != nil {
		return err
	}

	return nil
}
