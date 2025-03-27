package useCase

import (
	"app/pkg/security"
	u "app/pkg/utils"
	"app/userAuth"
	"fmt"
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

	stringJWT, err := userAuth.GenerateJWT(u.ToString(UserDB.Id), UserDB.Email, UserDB.Role)
	if err != nil {
		fmt.Println("Erro: ", err)
		return "", err
	}

	err = userAuth.SetUserSession(userAuth.UserSession{
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
	err := userAuth.LogoutUserSession(id)
	if err != nil {
		return err
	}
	return nil
}
