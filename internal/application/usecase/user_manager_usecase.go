// recebe a requisição já tratada
package usecase

import (
	"app/internal/application/dto"
	"app/internal/application/ports"
	"app/internal/domain/entity"
	"fmt"
)

type UserManagerUC struct {
	Repo    ports.IUserManager
	Service ports.IServices
}

func InitUserManager(repo ports.IUserManager, service ports.IServices) *UserManagerUC {
	return &UserManagerUC{repo, service}
}

func (it *UserManagerUC) RegisterUser(info dto.UserRegisterReq) error {
	hash, err := it.Service.HashPassword(info.Password)
	if err != nil {
		return fmt.Errorf("Error Bycript HashPassword")
	}

	newUser, err := entity.NewUser(
		it.Service.GenerateUUID(),
		info.Name,
		info.Email,
		hash,
	)

	if err != nil {
		return err
	}

	err = it.Repo.CreateUser(*newUser)
	if err != nil {
		return err
	}

	return nil
}

func (it *UserManagerUC) GetUser(infoID string) (*dto.UserRes, error) {
	result, err := it.Repo.GetUser(infoID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (it *UserManagerUC) GetUserList() (*[]dto.UserRes, error) {
	result, err := it.Repo.GetUserList()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (it *UserManagerUC) EditUser(id string, info dto.EditUserReq) error {
	if info.Password != nil {
		hash, err := it.Service.HashPassword(*info.Password)
		if err != nil {
			return fmt.Errorf("Error Bycript HashPassword")
		}
		info.Password = &hash
	}

	err := it.Repo.EditUser(id, info)

	if err != nil {
		return err
	}
	return nil
}

func (it *UserManagerUC) DeleteUser(infoID string) error {
	var err = it.Repo.DeleteUser(infoID)
	if err != nil {
		return err
	}
	return nil
}
