package model

import "fmt"

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

// Validate checks if the User fields are valid.
func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	if u.Password == "" {
		return fmt.Errorf("password is required")
	}
	if u.Role == "" {
		return fmt.Errorf("role is required")
	}
	return nil
}
