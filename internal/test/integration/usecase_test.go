package integration

import (
	"app/internal/model"
	"app/internal/repository"
	"app/internal/usecase"
	"app/internal/userauth"
	"database/sql"
	"strconv"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CreateToDelete(t *testing.T) {
	// layers
	conn, err := sql.Open("sqlite3", "dbtest.db")
	assert.NoError(t, err, err)
	repo := repository.NewSQLManager(conn)
	err = repo.CreateUserTable()
	assert.NoError(t, err, err)
	service := usecase.NewUserUseCase(repo)

	// test service.methods
	tempoAtual := time.Now().Format("0000-00-00 00:00:00")
	user := &model.User{
		Id:       0,
		Name:     "bellon",
		Email:    "bellon@gmail.com",
		Password: "bellon321",
		Role:     "user",
		Date:     tempoAtual,
	}

	err = user.Validate()
	require.NoError(t, err, err)

	err = service.UserCreate(*user)
	require.NoError(t, err, err)

	login := model.LoginFields{Email: user.Email, Password: user.Password}
	errList, err := login.ValidateFields()
	require.NoError(t, err, errList)

	keyJWT, err := service.UserLogin(login.Email, login.Password)
	require.NoError(t, err, err)

	_, claims := userauth.ValidateJWT(keyJWT)

	nId, err := strconv.Atoi(claims.UserID)
	err = service.UserDelete(nId)
	require.NoError(t, err, err)
}
