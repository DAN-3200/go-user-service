package middlewares

import (
	"app/internal/infrastructure/persistence/cache"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRole(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, err := cache.Session.GetInfoSession(ctx, "user_session")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Erro ao requerir dados de sessão")
			return
		}

		if role.Role != requiredRole {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Role inválida")
			return
		}
		ctx.Next()
	}
}
