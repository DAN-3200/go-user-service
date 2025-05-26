package server

import (
	"app/internal/controller"
	"app/internal/db"
	"app/internal/middlewares"
	"app/internal/repository"
	"app/internal/usecase"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	var server = gin.Default()
	// SQL Database Connection
	var ConnDB = db.Conn_Sqlite()
	defer ConnDB.Close()

	var dbManager = repository.NewSQLManager(ConnDB)
	// dbManager.CreateUserTable()

	var setUseCase = usecase.NewUserUseCase(dbManager)
	var setController = controller.NewUserController(*setUseCase)
	UserRoutes(server, *setController)

	server.Run(":3000")
}

func UserRoutes(server *gin.Engine, useController controller.UserController) {
	server.POST("/createUser", useController.CreateUser)
	server.GET("/readUser/:id", middlewares.AuthJWT(false), useController.ReadUser)
	server.GET("/readAllUser", middlewares.AuthJWT(true), useController.ReadAllUser)
	server.PUT("/updateUser", middlewares.AuthJWT(true), useController.UpdateUser)
	server.DELETE("/deleteUser/:id", middlewares.AuthJWT(true), useController.DeleteUser)

	// Login
	server.POST("/loginUser", useController.LoginUser)
	server.POST("/logoutUser", middlewares.AuthJWT(true), useController.LogoutUser)
}
