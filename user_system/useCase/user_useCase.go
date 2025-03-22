// recebe a requisição já tratada
package useCase

import (
	"app/contracts"
	"app/model"
	"fmt"
)

type UserUseCase struct {
	db contracts.IDB
}

func NewUserUseCase(db contracts.IDB) *UserUseCase {
	return &UserUseCase{db}
}

// -- Methods

func (it *UserUseCase) UserRead() []model.User {
	var result, err = it.db.UserReadSQL()
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return []model.User{}
	}

	return result
}

func (it *UserUseCase) UserCreate(info model.User) error {
	var err = it.db.UserSaveSQL(info)
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return err
	}
	return nil
}
