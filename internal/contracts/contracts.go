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
}
