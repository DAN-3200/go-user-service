package usecase

import (
	"app/internal/domain/dto"
	"app/internal/domain/entity"
	"fmt"
	"time"
)

func (it *LayerUseCase) LoginUser(info dto.Login) (string, error) {
	UserDB, err := it.Repo.LoginUserSQL(info.Email)
	if err != nil {
		return "", err
	}

	err = it.Service.CompareHashPassword(UserDB.PasswordHash, info.Password)
	if err != nil {
		return "", err
	}

	stringJWT, err := it.Service.GenerateJWT(UserDB.ID, UserDB.Role)
	if err != nil {
		return "", err
	}

	err = it.Service.SetUserSession(entity.UserSession{
		Id:    UserDB.ID,
		Name:  UserDB.Name,
		Email: UserDB.Email,
		Role:  UserDB.Role,
		JWT:   stringJWT,
	})

	if err != nil {
		return "", err
	}

	return stringJWT, nil
}

func (it *LayerUseCase) LogoutUser(infoID string) error {
	err := it.Service.LogoutUserSession(infoID)
	if err != nil {
		return err
	}
	return nil
}

func (it *LayerUseCase) RegisterUser(info dto.UserRegisterReq) error {
	hash, err := it.Service.HashPassword(info.Password)
	if err != nil {
		return fmt.Errorf("Error Bycript HashPassword")
	}

	newUser := entity.User{
		ID:              it.Service.GenerateUUID(),
		Name:            info.Name,
		Email:           info.Email,
		PasswordHash:    hash,
		IsEmailVerified: false,
		IsActive:        false,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Role:            "user",
	}

	err = it.Repo.CreateUserSQL(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (it *LayerUseCase) SendRefreshForEmail(email string) error {
	userDB, err := it.Repo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	stringJWT, err := it.Service.GenerateJWT(userDB.ID, userDB.Role)
	if err != nil {
		return err
	}

	err = it.Service.SendMail(email,
		"Codigo para redefinir senha:\n"+stringJWT,
	)
	if err != nil {
		return err
	}

	return nil
}

func (it *LayerUseCase) RefreshPassword(id string, info dto.RefreshPassword) error {
	hash, err := it.Service.HashPassword(info.NewPassword)
	if err != nil {
		return err
	}

	info.NewPassword = hash
	err = it.Repo.RefreshPassword(id, info)
	if err != nil {
		return err
	}

	return nil
}

func (it *LayerUseCase) ValidateEmail(email string) error {
	err := it.Repo.ValidateEmail(email)
	if err != nil {
		return err
	}
	return nil
}
