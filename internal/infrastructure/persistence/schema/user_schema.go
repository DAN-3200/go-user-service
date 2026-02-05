package schema

import (
	"app/internal/domain/entity"
	"app/internal/domain/vo"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID              string    `json:"id"` // UUID
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	PasswordHash    string    `json:"-"`
	IsEmailVerified bool      `json:"isEmailVerified"`
	IsActive        bool      `json:"isActive"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Role            string    `json:"role"`
}

func CreateUserTable(db *sql.DB) error {
	var _, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			is_email_verified BOOLEAN DEFAULT FALSE,
			is_active BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			role VARCHAR(50) NOT NULL
		);`,
	)

	if err != nil {
		fmt.Println("Erro:", err)
		return err
	}

	return nil
}

func ReconstituteUser(s *User) (*entity.User, error) {

	if s.ID == "" {
		return nil, errors.New("invalid user id")
	}

	return entity.UserDataFull(
		s.ID,
		s.Name,
		s.Email,
		s.PasswordHash,
		s.IsEmailVerified,
		s.IsActive,
		s.CreatedAt,
		s.UpdatedAt,
		vo.Role(s.Role),
	), nil
}
