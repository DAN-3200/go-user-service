package usecase

import (
	"app/internal/dto"
	"app/internal/model"
)

func (it *LayerUseCase) GetMyInfo(infoID string) (model.User, error) {
	myInfo, err := it.Repo.GetMyInfoSQL(infoID)
	if err != nil {
		return model.User{}, err
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
