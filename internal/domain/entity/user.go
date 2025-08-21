package entity

import "time"

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
