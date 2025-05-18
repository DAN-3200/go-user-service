package integration

import (
	"app/internal/db"
	"app/internal/model"
	"app/internal/repository"
	"app/internal/usecase"
	"app/internal/userauth"
	"testing"
	"time"
)

func Test_UserStartToEnd(t *testing.T) {
	ConnDB := db.Conn_Sqlite()
	dbManager := repository.NewSQLManager(ConnDB)
	actions := usecase.NewUserUseCase(dbManager)

	tempoAtual := time.Now().Format("2006-01-02 15:04:05")
	User := model.NewUser(
		"person",
		"person@gmail.com",
		"person321",
		"user",
		tempoAtual,
	)

	actions.UserCreate(*User)
	resJwt, err := actions.UserLogin(User.Email, User.Password)
	if err != nil {
		t.Fatal("Login não efetuado")
	}
	ok, claims := userauth.ValidateJWT(resJwt)
	if !ok {
		t.Fatal("JWT não validado")
	}
	actions.UserLogout(claims.UserID)
}
