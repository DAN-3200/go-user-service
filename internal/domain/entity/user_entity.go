package entity

import (
	"app/internal/domain/vo"
	"errors"
	"strings"
	"time"
)

type User struct {
	_ID              string
	_Name            string
	_Email           string
	_PasswordHash    string
	_IsEmailVerified bool
	_IsActive        bool
	_CreatedAt       time.Time
	_UpdatedAt       time.Time
	_Role            vo.Role
}

func NewUser(id string, name string, email string, passwordHash string) (*User, error) {

	if id == "" {
		return nil, errors.New("id required")
	}

	if len(strings.TrimSpace(name)) < 3 {
		return nil, errors.New("name invalid")
	}

	if email == "" {
		return nil, errors.New("email invalid")
	}

	if passwordHash == "" {
		return nil, errors.New("password invalid")
	}

	return &User{
		_ID:              id,
		_Name:            strings.ToLower(name),
		_Email:           email,
		_PasswordHash:    passwordHash,
		_IsEmailVerified: false,
		_IsActive:        false,
		_CreatedAt:       time.Now(),
		_UpdatedAt:       time.Now(),
		_Role:            vo.RoleUser,
	}, nil
}

func UserDataFull(
	id string,
	name string,
	email string,
	passwordHash string,
	isEmailVerified bool,
	isActive bool,
	createdAt time.Time,
	updatedAt time.Time,
	role vo.Role,
) *User {

	return &User{
		_ID:              id,
		_Name:            name,
		_Email:           email,
		_PasswordHash:    passwordHash,
		_IsEmailVerified: isEmailVerified,
		_IsActive:        isActive,
		_CreatedAt:       createdAt,
		_UpdatedAt:       updatedAt,
		_Role:            role,
	}
}


// getters

func (it *User) ID() string {
	return it._ID
}

func (it *User) Name() string {
	return it._Name
}

func (it *User) Id() string {
	return it._ID
}

func (it *User) Email() string {
	return it._Email
}

func (it *User) PasswordHash() string {
	return it._PasswordHash
}

func (it *User) IsEmailVerified() bool {
	return it._IsEmailVerified
}

func (it *User) IsActive() bool {
	return it._IsActive
}

func (it *User) CreatedAt() time.Time {
	return it._CreatedAt
}

func (it *User) UpdatedAt() time.Time {
	return it._UpdatedAt
}

func (it *User) Role() vo.Role {
	return it._Role
}

// setters

func (u *User) SetEmail(email string) {
	u._Email = email
}

func (it *User) SetName(name string) {
	it._Name = name
}

func (u *User) SetPasswordHash(passwordHash string) {
	u._PasswordHash = passwordHash
}

func (u *User) SetIsEmailVerified(isVerified bool) {
	u._IsEmailVerified = isVerified
}

func (u *User) SetIsActive(isActive bool) {
	u._IsActive = isActive
}

func (u *User) SetRole(role vo.Role) {
	u._Role = role
}
