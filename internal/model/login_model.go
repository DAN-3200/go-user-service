package model

import (
	"fmt"
	"strings"
)

type LoginFields struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (it *LoginFields) Validate() error {
	if it.Email == "" {
		return fmt.Errorf("email is required")
	}
	if !strings.HasSuffix(it.Email, "@gmail.com") && !strings.HasSuffix(it.Email, "@hotmail.com") && !strings.HasSuffix(it.Email, "@outlook.com") {
		return fmt.Errorf("email invalid")
	}
	if it.Password == "" || len(it.Password) > 20 {
		return fmt.Errorf("password is invalid")
	}
	return nil
}

func (it *LoginFields) ValidateFields() ([]string, error) {
	var errs []string
	if it.Email == "" {
		errs = append(errs, "email is required")
	}
	if !strings.HasSuffix(it.Email, "@gmail.com") && !strings.HasSuffix(it.Email, "@hotmail.com") && !strings.HasSuffix(it.Email, "@outlook.com") {
		errs = append(errs, "email invalid")
	}
	if it.Password == "" || len(it.Password) > 20 || len(it.Password) < 5 {
		errs = append(errs, "password is invalid")
	}
	if len(errs) != 0 {
		return errs, fmt.Errorf("Error de validação")
	}
	return nil, nil
}
