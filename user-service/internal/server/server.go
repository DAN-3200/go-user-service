package server

import (
	"app/internal/controller"
	"app/internal/db"
	"app/internal/repository"
	"app/internal/useCase"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	var server = gin.Default()
	// SQL Database Connection
	var ConnDB = db.Conn_Sqlite()
	defer ConnDB.Close()

	var dbManager = repository.NewSQLManager(ConnDB)
	// dbManager.CreateUserTable()

	var setUseCase = useCase.NewUserUseCase(dbManager)
	var setController = controller.NewUserController(*setUseCase)
	_UserRoutes(server, *setController)

	server.Run(":3000")
}

func _UserRoutes(server *gin.Engine, useController controller.UserController) {
	server.GET("/readUser/:id", useController.ReadUser)
	server.GET("/readAllUser", useController.ReadAllUser)
	server.POST("/createUser", useController.CreateUser)
	server.PUT("/updateUser", useController.UpdateUser)
	server.DELETE("/deleteUser", useController.DeleteUser)

	// Login
	server.POST("/loginUser", useController.LoginUser)
	server.POST("/logoutUser", useController.LogoutUser)
}
