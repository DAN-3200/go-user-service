package repository

import (
	"app/internal/application/dto"
	"app/internal/domain/entity"
	"app/internal/infrastructure/persistence/schema"
	"app/pkg/utils"
	"database/sql"
	"fmt"
	"strings"
)

// ------------------------------------------------------------------------

type UserInfoRepo struct {
	DB *sql.DB
}

func NewUserInfoRepo(db *sql.DB) *UserInfoRepo {
	return &UserInfoRepo{db}
}

func (it *UserInfoRepo) GetMyInfo(infoID string) (*entity.User, error) {
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

	var obj schema.User
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
		return nil, err
	}
	
	res, err := schema.ReconstituteUser(&obj)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (it *UserInfoRepo) EditMyInfo(id string, info dto.EditMeReq) error {
	cols, args, err := utils.MapSQLInsertFields(
		map[string]string{
			"Name":     "name",
			"IsActive": "is_active",
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
