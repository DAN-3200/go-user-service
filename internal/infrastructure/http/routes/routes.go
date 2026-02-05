package routes

import (
	"app/internal/infrastructure/http/controller"
	mdw "app/internal/infrastructure/http/middlewares"
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SetUserManagerRoutes(server *gin.Engine, controller *controller.UserManagerController) {
	admin := server.Group("/users", mdw.Auth2FA(), mdw.AuthRole("admin"))
	{
		admin.POST("", controller.CreateUser)
		admin.GET("", controller.GetUserList)
		admin.GET(":id", controller.GetUser)
		admin.PATCH(":id", controller.EditUser)
		admin.DELETE(":id", controller.DeleteUser)
	}
}

func SetUserAuthRoutes(server *gin.Engine, controllerAuth *controller.UserAuthController, controllerCRUD *controller.UserManagerController) {
	auth := server.Group("/auth")
	{
		auth.POST("/login", controllerAuth.LoginUser)
		auth.POST("/logout", mdw.Auth2FA(), controllerAuth.LogoutUser)
		auth.POST("/register", controllerCRUD.CreateUser)
		auth.POST("/refresh-token")
		auth.POST("/verify-email")
		forgetPassword := auth.Group("/forget-password")
		{
			forgetPassword.GET("/send-token/:email", controllerAuth.SendRefreshForEmail)
			forgetPassword.POST("/refresh-password", controllerAuth.RefreshPassword)
		}
	}
}

func SetUserInfoRoutes(server *gin.Engine, controller *controller.UserInfoController) {
	me := server.Group("/me", mdw.Auth2FA())
	{
		me.GET("", controller.GetMyInfo)
		me.PATCH("", controller.EditMyInfo)
	}
}

func HealthCheck(server *gin.Engine, dbSQL *sql.DB, Redis *redis.Client) {
	server.GET("/health", func(ctx *gin.Context) {
		errList := map[string]any{
			"db-sql": true,
			"redis":  true,
		}
		if err := dbSQL.Ping(); err != nil {
			errList["db-sql"] = false
		}
		if err := Redis.Ping(context.Background()).Err(); err != nil {
			errList["redis"] = false
		}
		status := 200
		for _, y := range errList {
			if y == false {
				status = 500
			}
		}

		ctx.JSON(status, errList)
	})
}
