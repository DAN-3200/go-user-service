package usecase

import (
	"app/internal/application/dto"
	"app/internal/application/ports"

)

type UserInfoUC struct {
	Repo    ports.IUserMeInfo
	Service ports.IServices
}

func InitUserInfo(repo ports.IUserMeInfo, service ports.IServices) *UserInfoUC {
	return &UserInfoUC{repo, service}
}

func (it *UserInfoUC) GetMyInfo(infoID string) (*dto.UserInfoRes, error) {
	myInfo, err := it.Repo.GetMyInfo(infoID)
	if err != nil {
		return nil, err
	}
	return dto.ToUserResponse(myInfo), nil
}

func (it *UserInfoUC) EditMyInfo(id string, info dto.EditMeReq) error {
	err := it.Repo.EditMyInfo(id, info)
	if err != nil {
		return err
	}
	return nil
}
