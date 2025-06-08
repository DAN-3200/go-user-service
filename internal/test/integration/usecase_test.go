package integration

import (
	"app/internal/db"
	"app/internal/dto"
	"app/internal/repository"
	"app/internal/usecase"
	"app/internal/userauth"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CreateToDelete(t *testing.T) {
	// layers
	conn, err := sql.Open("sqlite3", ":memory:")
	userauth.InitCoreRedis(db.Conn_Redis())
	assert.NoError(t, err, err)
	repo := repository.NewSQLManager(conn)
	err = repo.CreateUserTable()
	require.NoError(t, err, err)
	service := usecase.NewUserUseCase(repo)

	// test service.methods
	user := &dto.UserReq{
		Name:     "bellon",
		Email:    "bellon@gmail.com",
		Password: "bellon321",
		Role:     "user",
	}


	err = service.CreateUser(*user)
	require.NoError(t, err, err)

	login := dto.Login{Email: user.Email, Password: user.Password}


	keyJWT, err := service.UserLogin(login)
	require.NoError(t, err, err)

	_, claims := userauth.ValidateJWT(keyJWT)

	err = service.DeleteUser(claims.UserID)
	require.NoError(t, err, err)
}
