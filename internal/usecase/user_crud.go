// recebe a requisição já tratada
package usecase

import (
	"app/internal/contracts"
	"app/internal/dto"
	"app/internal/model"
	"app/pkg/security"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserUseCase struct {
	Repo contracts.UserRepoSQL
}

func NewUserUseCase(repo contracts.UserRepoSQL) *UserUseCase {
	return &UserUseCase{repo}
}

// ------------------------------------------------------------------------

func (it *UserUseCase) CreateUser(info dto.UserReq) error {
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

func (it *UserUseCase) GetUser(infoID string) (dto.UserRes, error) {
	result, err := it.Repo.UserReadSQL(infoID)
	if err != nil {
		return dto.UserRes{}, err
	}

	return result, nil
}

func (it *UserUseCase) GetAllUsers() ([]dto.UserRes, error) {
	result, err := it.Repo.ReadAllUserSQL()
	if err != nil {
		return []dto.UserRes{}, err
	}

	return result, nil
}

func (it *UserUseCase) UpdateUser(info dto.UserUpdateReq) error {
	var err = it.Repo.UserUpdateSQL(info)
	if err != nil {
		return err
	}
	return nil
}

func (it *UserUseCase) DeleteUser(infoID string) error {
	var err = it.Repo.UserDeleteSQL(infoID)
	if err != nil {
		return err
	}
	return nil
}
