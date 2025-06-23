package contracts

import (
	"app/internal/dto"
	"app/internal/model"
)

// Contrato para Bancos SQL
type UserRepoSQL interface {
	CreateUserSQL(info model.User) error
	GetUserSQL(infoID string) (dto.UserRes, error)
	GetUserListSQL() ([]dto.UserRes, error)
	EditUserSQL(id string, info dto.EditUserReq) error
	DeleteUserSQL(infoID string) error
	//
	LoginUserSQL(UserEmail string) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	RefreshPassword(id string, info dto.RefreshPassword) error
	ValidateEmail(email string) error
	//
	GetMyInfoSQL(infoID string) (model.User, error)
	EditMyInfoSQL(id string, info dto.EditMeReq) error
}

type Drive interface {
	SendMail(to string, body string) error
}
