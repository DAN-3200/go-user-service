package dto

import (
	"app/internal/mytypes"
	"strings"
)

type UserReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5,max=20"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
}

type UserRes struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"createAt"`
}

type UserUpdateReq struct {
	ID           string `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required,min=5,max=20"`
	PasswordHash string `json:"password" binding:"required"`
}

func (it *UserReq) ValidateFields() error {
	errs := mytypes.ErrorsList{}

	if it.Name == "" || len(it.Name) > 20 || len(it.Name) < 4 {
		errs = append(errs, "name is invalid")
	}

	if it.Email == "" {
		errs = append(errs, "email is required")
	}

	if !strings.HasSuffix(it.Email, "@gmail.com") && !strings.HasSuffix(it.Email, "@hotmail.com") && !strings.HasSuffix(it.Email, "@outlook.com") {
		errs = append(errs, "email invalid")
	}

	if it.Password == "" || len(it.Password) > 20 || len(it.Password) < 5 {
		errs = append(errs, "password is invalid")
	}

	if it.Role == "" {
		errs = append(errs, "role is required")
	}

	if it.Role != "user" && it.Role != "admin" {
		errs = append(errs, "role invalid")
	}

	if len(errs) != 0 {
		return errs
	}

	return nil
}
