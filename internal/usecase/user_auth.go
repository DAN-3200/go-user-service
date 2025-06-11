package usecase

import (
	"app/internal/dto"
	"app/internal/model"
	"app/internal/userauth"
	"app/pkg/security"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"github.com/google/uuid"
)

func (it *LayerUseCase) UserLogin(info dto.Login) (string, error) {
	UserDB, err := it.Repo.LoginUserSQL(info.Email)
	if err != nil {
		return "", err
	}

	err = security.CompareHashPassword(UserDB.PasswordHash, info.Password)
	if err != nil {
		return "", err
	}

	stringJWT, err := userauth.GenerateJWT(UserDB.ID, UserDB.Role)
	if err != nil {
		return "", err
	}

	err = userauth.SetUserSession(userauth.UserSession{
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

func (it *LayerUseCase) UserLogout(infoID string) error {
	err := userauth.LogoutUserSession(infoID)
	if err != nil {
		return err
	}
	return nil
}

func (it *LayerUseCase) UserRegister(info dto.UserRegisterReq) error {
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
		IsActive:        false,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Role:            "user",
	}

	err = it.Repo.UserSaveSQL(newUser)
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

	stringJWT, err := userauth.GenerateJWT(userDB.ID, userDB.Role)
	if err != nil {
		return err
	}

	// Simple Mail Transfer Protocol (smtp)
	from := os.Getenv("MYEMAIL")
	password := os.Getenv("MYPASSWORD")
	println(from, password)
	to := []string{email}
	msg := []byte("Codigo para redefinir senha:\n" + stringJWT)

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	err = smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		return err
	}

	return nil
}
