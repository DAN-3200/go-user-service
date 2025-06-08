package server

import (
	"app/internal/controller"
	"app/internal/db"
	"app/internal/repository"
	"app/internal/usecase"
	"app/internal/userauth"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	server := gin.Default()

	ConnSQL := db.Conn_Postgres()
	defer ConnSQL.Close()

	ConnRedis := db.Conn_Redis()
	userauth.InitCoreRedis(ConnRedis)
	defer ConnRedis.Close()

	dbManager := repository.NewSQLManager(ConnSQL)
	dbManager.CreateUserTable()

	HealthCheck(server, ConnSQL, ConnRedis)
	Routes(server,*controller.Init(*usecase.Init(dbManager)))

	server.Run(":3000")
}
