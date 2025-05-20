package model

import (
	"fmt"
	"strings"
)

type User struct {
	Id       int    `json:"id"` // temporáriamente
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Date     string `json:"date"`
}

func NewUser(name, email, password, role, date string) *User {
	return &User{Name: name, Email: email, Password: password, Role: role, Date: date}
}

func (u *User) Validate() error {
	if u.Name == "" || len(u.Name) > 20 {
		return fmt.Errorf("name is invalid")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	if !strings.HasSuffix(u.Email, "@gmail.com") && !strings.HasSuffix(u.Email, "@hotmail.com") && !strings.HasSuffix(u.Email, "@outlook.com") {
		return fmt.Errorf("email invalid")
	}
	if u.Password == "" || len(u.Password) > 20  {
		return fmt.Errorf("password is invalid")
	}
	if u.Role == "" {
		return fmt.Errorf("role is required")
	}
	if u.Role != "user" && u.Role != "admin" {
		return fmt.Errorf("role invalid")
	}
	return nil
}
