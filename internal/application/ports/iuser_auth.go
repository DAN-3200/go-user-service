package ports

import (
	"app/internal/application/dto"
	"app/internal/domain/entity"
)

type IAuthUser interface {
	LoginUser(UserEmail string) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	RefreshPassword(id string, info dto.RefreshPassword) error
	ValidateEmail(email string) error
}