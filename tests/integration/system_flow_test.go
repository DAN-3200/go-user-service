package integration

import (
	// "app/internal/inner/dto"
	// "app/internal/inner/usecase"
	// "app/internal/outer/adapters"
	// "app/internal/outer/persistence/cache"
	// "app/internal/outer/persistence/db"
	// "app/internal/outer/persistence/repository"
	// "app/internal/outer/persistence/schema"
	// "database/sql"
	"testing"

	// _ "github.com/mattn/go-sqlite3"

	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

func Test_CreateToDelete(t *testing.T) {
	// // layers
	// conn, err := sql.Open("sqlite3", ":memory:")
	// cache_db := db.Conn_Redis()
	// cache.InitCoreRedis(cache_db)
	// assert.NoError(t, err, err)
	// repo := repository.NewUserManagerRepo(conn)
	// err = schema.CreateUserTable(cache_db)
	// require.NoError(t, err, err)
	// service := usecase.InitUserManager(repo, adapters.Static)

	// // test service.methods
	// user := &dto.UserReq{
	// 	Name:     "bellon",
	// 	Email:    "bellon@gmail.com",
	// 	Password: "bellon321",
	// 	Role:     "user",
	// }

	// err = service.Repo.CreateUser(*user)
	// require.NoError(t, err, err)

	// login := dto.Login{Email: user.Email, Password: user.Password}

	// repoAuth := usecase.InitUserAuth()
	// keyJWT, err := service.LoginUser(login)
	// require.NoError(t, err, err)

	// _, claims := adapters.Static.ValidateJWT(keyJWT)

	// err = service.DeleteUser(claims.UserID)
	// require.NoError(t, err, err)
}
