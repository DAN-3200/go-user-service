package usecase

import (
	"app/internal/domain/dto"
	"app/internal/domain/entity"
)

func (it *LayerUseCase) GetMyInfo(infoID string) (entity.User, error) {
	myInfo, err := it.Repo.GetMyInfoSQL(infoID)
	if err != nil {
		return entity.User{}, err
	}
	return myInfo, nil
}

func (it *LayerUseCase) EditMyInfo(id string, info dto.EditMeReq) error {
	err := it.Repo.EditMyInfoSQL(id, info)
	if err != nil {
		return err
	}
	return nil
}
