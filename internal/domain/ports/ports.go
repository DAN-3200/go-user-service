package ports

import (
	"app/internal/domain/dto"
	"app/internal/domain/entity"
)

type RepositoryPorts interface {
	CreateUserSQL(info entity.User) error
	GetUserSQL(infoID string) (dto.UserRes, error)
	GetUserListSQL() ([]dto.UserRes, error)
	EditUserSQL(id string, info dto.EditUserReq) error
	DeleteUserSQL(infoID string) error
	//
	LoginUserSQL(UserEmail string) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	RefreshPassword(id string, info dto.RefreshPassword) error
	ValidateEmail(email string) error
	//
	GetMyInfoSQL(infoID string) (entity.User, error)
	EditMyInfoSQL(id string, info dto.EditMeReq) error
}

type ServicePorts interface {
	SendMail(to string, body string) error
	GenerateUUID() string
	HashPassword(password string) (string, error)
	CompareHashPassword(pivotPassword, inputPassword string) error
	SetUserSession(info entity.UserSession) error
	GetUserSession(Id string) (*entity.UserSession, error)
	LogoutUserSession(Id string) error
	GenerateJWT(userID, userRole string) (string, error)
}
