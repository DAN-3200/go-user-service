package usecase

import (
	"app/internal/model"
)

func (it *LayerUseCase) GetMyInfo(infoID string) (model.User, error) {
	myInfo, err := it.Repo.MyInfoSQL(infoID)
	if err != nil {
		return model.User{}, err
	}
	return myInfo, nil
}

func (it *LayerUseCase) EditMyInfo(info map[string]any) error {
	err := it.Repo.EditMyInfoSQL(info)
	if err != nil {
		return err
	}
	return nil
}
