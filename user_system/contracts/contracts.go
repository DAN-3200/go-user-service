package contracts

import "app/model"

// Contrato para Bancos SQL
type IDB interface {
	UserSaveSQL(info model.User) error
	UserReadSQL(infoID int) (model.User, error)
	ReadAllUserSQL() ([]model.User, error)
	UserUpdateSQL(info model.User) error
	UserDeleteSQL(infoID int) error
}
