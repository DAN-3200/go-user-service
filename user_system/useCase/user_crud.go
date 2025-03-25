// recebe a requisição já tratada
package useCase

import (
	"app/contracts"
	"app/model"
	"fmt"
)

type UserUseCase struct {
	DB contracts.IDB
}

func NewUserUseCase(db contracts.IDB) *UserUseCase {
	return &UserUseCase{db}
}

// -- Methods

func (it *UserUseCase) UserRead(infoID int) model.User {
	var result, err = it.DB.UserReadSQL(infoID)
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return model.User{}
	}

	return result
}

func (it *UserUseCase) ReadAllUser() ([]model.User, error) {
	var result, err = it.DB.ReadAllUserSQL()
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return []model.User{}, err
	}

	return result, nil
}

func (it *UserUseCase) UserCreate(info model.User) error {
	var err = it.DB.UserSaveSQL(info)
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return err
	}
	return nil
}

func (it *UserUseCase) UserUpdate(info model.User) error {
	var err = it.DB.UserUpdateSQL(info)
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return err
	}
	return nil
}

func (it *UserUseCase) UserDelete(idUser int) error {
	var err = it.DB.UserDeleteSQL(idUser)
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return err
	}
	return nil
}
