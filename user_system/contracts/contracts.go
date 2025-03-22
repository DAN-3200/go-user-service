package contracts

import "app/model"

// Contrato para Bancos SQL
type IDB interface {
	UserSaveSQL(info model.User) error
	UserReadSQL() ([]model.User, error)
	UserUpdateSQL(info model.User) error
	UserDeleteSQL(info struct{ Id string }) error
}
