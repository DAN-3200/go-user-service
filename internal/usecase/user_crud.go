// recebe a requisição já tratada
package usecase

import (
	"app/internal/dto"
	"app/internal/model"
	"app/pkg/security"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (it *LayerUseCase) CreateUser(info dto.UserReq) error {
	hash, err := security.HashPassword(info.Password)
	if err != nil {
		return fmt.Errorf("Error Bycript HashPassword")
	}

	newUser := model.User{
		ID:              uuid.New().String(),
		Name:            info.Name,
		Email:           info.Email,
		PasswordHash:    hash,
		IsEmailVerified: false,
		IsActive:        true,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Role:            info.Role,
	}

	err = it.Repo.UserSaveSQL(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (it *LayerUseCase) GetUser(infoID string) (dto.UserRes, error) {
	result, err := it.Repo.UserReadSQL(infoID)
	if err != nil {
		return dto.UserRes{}, err
	}

	return result, nil
}

func (it *LayerUseCase) GetAllUsers() ([]dto.UserRes, error) {
	result, err := it.Repo.ReadAllUserSQL()
	if err != nil {
		return []dto.UserRes{}, err
	}

	return result, nil
}

func (it *LayerUseCase) UpdateUser(info dto.UserUpdateReq) error {
	var err = it.Repo.UserUpdateSQL(info)
	if err != nil {
		return err
	}
	return nil
}

func (it *LayerUseCase) DeleteUser(infoID string) error {
	var err = it.Repo.UserDeleteSQL(infoID)
	if err != nil {
		return err
	}
	return nil
}
