package ports

import (
	"app/internal/application/dto"
	"app/internal/domain/entity"
)

type IUserMeInfo interface {
	GetMyInfo(infoID string) (*entity.User, error)
	EditMyInfo(id string, info dto.EditMeReq) error
}

