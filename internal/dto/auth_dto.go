package dto

import (
	"app/internal/mytypes"
	"strings"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (it *Login) ValidateFields() error {
	errs := mytypes.ErrorsList{}

	if it.Email == "" {
		errs = append(errs, "email is required")
	}
	if !strings.HasSuffix(it.Email, "@gmail.com") && !strings.HasSuffix(it.Email, "@hotmail.com") && !strings.HasSuffix(it.Email, "@outlook.com") {
		errs = append(errs, "email invalid")
	}
	if it.Password == "" || len(it.Password) > 20 || len(it.Password) < 5 {
		errs = append(errs, "password invalid")
	}
	if len(errs) != 0 {
		return errs
	}

	return nil
}

type UserRegisterRes struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
