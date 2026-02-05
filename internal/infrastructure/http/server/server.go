package server

import (
	"app/internal/application/usecase"
	"app/internal/infrastructure/adapters"
	"app/internal/infrastructure/http/controller"
	"app/internal/infrastructure/http/routes"
	"app/internal/infrastructure/persistence/cache"
	"app/internal/infrastructure/persistence/db"
	"app/internal/infrastructure/persistence/repository"
	"app/internal/infrastructure/persistence/schema"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	server := gin.Default()

	connRedis := db.Conn_Redis()
	defer connRedis.Close()
	
	cache.InitCoreRedis(connRedis)

	connPostgres := db.Conn_Postgres()
	defer connPostgres.Close()

	schema.CreateUserTable(connPostgres)

	userManager := repository.NewUserManagerRepo(connPostgres)
	userAuth := repository.NewUserAuthRepo(connPostgres)
	userInfo := repository.NewUserInfoRepo(connPostgres)

	newServices := adapters.NewInstanceService()

	routes.HealthCheck(server, connPostgres, connRedis)

	userAuthHandler := controller.InitUserAuth(
		usecase.InitUserAuth(userAuth, newServices, cache.Session),
	)

	userInfoHandler := controller.InitUserInfo(
		usecase.InitUserInfo(userInfo, newServices),
	)

	userManagerHandler := controller.InitUserManager(
		usecase.InitUserManager(userManager, newServices),
	)

	routes.SetUserAuthRoutes(server,
		userAuthHandler,
		userManagerHandler,
	)

	routes.SetUserInfoRoutes(server,
		userInfoHandler,
	)

	routes.SetUserManagerRoutes(server,
		userManagerHandler,
	)

	server.Run(":3000")
}
