package server

import (
	"app/internal/controller"
	mdw "app/internal/middlewares"
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Routes(server *gin.Engine, controller controller.UserController) {
	admin := server.Group("/users", mdw.AuthJWT())
	{
		admin.POST("", controller.CreateUser)
		admin.GET("", controller.GetAllUsers)
		admin.GET(":id", controller.GetUser)
		admin.PUT(":id", controller.UpdateUser)
		admin.DELETE(":id", controller.DeleteUser)
	}

	auth := server.Group("/auth")
	{
		auth.POST("/login", controller.LoginUser)
		auth.POST("/logout", mdw.AuthJWT(), controller.LogoutUser)
		auth.POST("/register", controller.RegisterUser)
		auth.POST("/refresh", func(ctx *gin.Context) {
			
		})
		auth.POST("/forgot-password")
	}

	me := server.Group("/me", mdw.AuthJWT())
	{
		me.GET("", controller.MyInfo)
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
