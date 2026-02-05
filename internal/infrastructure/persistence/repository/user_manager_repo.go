package repository

import (
	"app/internal/application/dto"
	"app/internal/domain/entity"
	"app/pkg/utils"
	"database/sql"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// ------------------------------------------------------------------------
type UserManagerRepo struct {
	DB *sql.DB
}

func NewUserManagerRepo(db *sql.DB) *UserManagerRepo {
	return &UserManagerRepo{db}
}

func (it *UserManagerRepo) CreateUser(info entity.User) error {
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
		info.ID(),
		info.Name(),
		info.Email(),
		info.PasswordHash(),
		info.IsEmailVerified(),
		info.IsActive(),
		info.CreatedAt(),
		info.UpdatedAt(),
		info.Role(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (it *UserManagerRepo) GetUser(infoID string) (*dto.UserRes, error) {
	query, args, err := sq.Select("id", "name", "email", "role", "created_at").From("users").Where(sq.Eq{"id": infoID}).ToSql()
	if err != nil {
		return nil, err
	}
	row := it.DB.QueryRow(query, args...)

	var userObj dto.UserRes
	err = row.Scan(
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
		return nil, err
	}

	return &userObj, nil
}

func (it *UserManagerRepo) GetUserList() (*[]dto.UserRes, error) {
	query := `SELECT id, name, email, role, created_at FROM users`
	rows, err := it.DB.Query(query)
	if err != nil {
		fmt.Printf("Erro de consulta: %v", err)
		return nil, err
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
			return nil, err
		}
		userList = append(userList, m)
	}
	rows.Close()
	return &userList, nil
}

func (it *UserManagerRepo) EditUser(id string, info dto.EditUserReq) error {
	cols, args, err := utils.MapSQLInsertFields(
		map[string]string{
			"Name":            "name",
			"Email":           "email",
			"Password":        "password_hash", // hash armazenado
			"IsEmailVerified": "is_email_verified",
			"IsActive":        "is_active",
			"Role":            "role",
		},
		info,
	)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`UPDATE users SET %s WHERE id='%s'`, strings.Join(cols, ", "), id)
	_, err = it.DB.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (it *UserManagerRepo) DeleteUser(infoID string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := it.DB.Exec(query, infoID)
	if err != nil {
		fmt.Println("Erro ao excluir: ", err)
		return err
	}
	return nil
}
