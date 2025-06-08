package usecase

import "app/internal/model"

func (it *LayerUseCase) MyInfo(infoID string) (model.User, error) {
	myInfo, err := it.Repo.MyInfoSQL(infoID)
	if err != nil {
		return model.User{}, err
	}
	return myInfo, nil
}
