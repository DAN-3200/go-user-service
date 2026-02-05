package ports

import (
	"app/internal/application/dto"
	"app/internal/domain/entity"
)


type IUserManager interface {
	CreateUser(info entity.User) error
	GetUser(id string) (*dto.UserRes, error)
	GetUserList() (*[]dto.UserRes, error)
	EditUser(id string, info dto.EditUserReq) error
	DeleteUser(id string) error
}
