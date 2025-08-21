package server

import (
	"app/internal/domain/usecase"
	"app/internal/infrastructure/adapters"
	"app/internal/infrastructure/db"
	"app/internal/infrastructure/http/controller"
	"app/internal/infrastructure/http/routes"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	server := gin.Default()

	ConnSQL := db.Conn_Postgres()
	defer ConnSQL.Close()

	ConnRedis := db.Conn_Redis()
	adapters.InitCoreRedis(ConnRedis)
	defer ConnRedis.Close()

	dbManager := adapters.NewSQLManager(ConnSQL)
	dbManager.CreateUserTable()

	routes.HealthCheck(server, ConnSQL, ConnRedis)
	// middlewares.SetProme(server)
	routes.SetRoutes(server,
		controller.Init(
			usecase.Init(dbManager, adapters.LayerService()),
		),
	)

	server.Run(":3000")
}
