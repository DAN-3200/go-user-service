package contracts

import (
	"app/internal/dto"
	"app/internal/model"
)

// Contrato para Bancos SQL
type UserRepoSQL interface {
	UserSaveSQL(info model.User) error
	UserReadSQL(infoID string) (dto.UserRes, error)
	ReadAllUserSQL() ([]dto.UserRes, error)
	UserUpdateSQL(info dto.UserUpdateReq) error
	UserDeleteSQL(infoID string) error
	LoginUserSQL(UserEmail string) (model.User, error)
	MyInfoSQL(infoID string) (model.User, error)
	EditMyInfoSQL(info map[string]any) error
	GetUserByEmail(email string) (model.User, error)
	RefreshPassword(info dto.RefreshPassword) error
	ValidateEmail(email string) error
}

type Drive interface {
	SendMail(to string, body string) error
}
