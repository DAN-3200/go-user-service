package integration

import (
	"app/internal/domain/dto"
	"app/internal/domain/usecase"
	"app/internal/infrastructure/adapters"
	"app/internal/infrastructure/db"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CreateToDelete(t *testing.T) {
	// layers
	conn, err := sql.Open("sqlite3", ":memory:")
	adapters.InitCoreRedis(db.Conn_Redis())
	assert.NoError(t, err, err)
	repo := adapters.NewSQLManager(conn)
	err = repo.CreateUserTable()
	require.NoError(t, err, err)
	service := usecase.Init(repo, adapters.LayerService())

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

	keyJWT, err := service.LoginUser(login)
	require.NoError(t, err, err)

	_, claims := adapters.Static.ValidateJWT(keyJWT)

	err = service.DeleteUser(claims.UserID)
	require.NoError(t, err, err)
}
