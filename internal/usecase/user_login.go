package usecase

import (
	"app/internal/userauth"
	"app/pkg/security"
	"fmt"
	"strconv"
)

func (it *UserUseCase) UserLogin(UserEmail string, UserPassword string) (string, error) {
	UserDB, err := it.DB.LoginUserSQL(UserEmail)
	if err != nil {
		return "", err
	}

	err = security.CompareHashPassword(UserDB.Password, UserPassword)
	if err != nil {
		return "", err
	}

	stringJWT, err := userauth.GenerateJWT(strconv.Itoa(UserDB.Id), UserDB.Email, UserDB.Role)
	if err != nil {
		fmt.Println("Erro: ", err)
		return "", err
	}

	err = userauth.SetUserSession(userauth.UserSession{
		Id:    UserDB.Id,
		Name:  UserDB.Name,
		Email: UserDB.Email,
		Role:  UserDB.Role,
		JWT:   stringJWT,
	})
	
	if err != nil {
		fmt.Println("Erro: ", err)
		return "", err
	}

	return stringJWT, nil
}

func (it *UserUseCase) UserLogout(id string) error {
	err := userauth.LogoutUserSession(id)
	if err != nil {
		return err
	}
	return nil
}
