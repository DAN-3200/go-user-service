package usecase

import (
	"app/internal/application/dto"
	"app/internal/application/ports"
)

type UserAuthUC struct {
	Repo    ports.IAuthUser
	Service ports.IServices
	Session ports.ISession
}

func InitUserAuth(repo ports.IAuthUser, service ports.IServices, session ports.ISession) *UserAuthUC {
	return &UserAuthUC{repo, service, session}
}

func (it *UserAuthUC) LoginUser(info dto.Login) (string, error) {
	UserDB, err := it.Repo.LoginUser(info.Email)
	if err != nil {
		return "", err
	}

	err = it.Service.CompareHashPassword(UserDB.PasswordHash(), info.Password)
	if err != nil {
		return "", err
	}

	stringJWT, err := it.Service.GenerateJWT(UserDB.ID(), string(UserDB.Role()))
	if err != nil {
		return "", err
	}

	err = it.Session.SetUserSession(
		UserDB.ID(),
		UserDB.Name(),
		UserDB.Email(),
		string(UserDB.Role()),
		stringJWT,
	)

	if err != nil {
		return "", err
	}

	return stringJWT, nil
}

func (it *UserAuthUC) LogoutUser(infoID string) error {
	err := it.Session.LogoutUserSession(infoID)
	if err != nil {
		return err
	}
	return nil
}

func (it *UserAuthUC) SendRefreshForEmail(email string) error {
	userDB, err := it.Repo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	stringJWT, err := it.Service.GenerateJWT(userDB.ID(), string(userDB.Role()))
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

func (it *UserAuthUC) RefreshPassword(id string, info dto.RefreshPassword) error {
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

func (it *UserAuthUC) ValidateEmail(email string) error {
	err := it.Repo.ValidateEmail(email)
	if err != nil {
		return err
	}
	return nil
}
