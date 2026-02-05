package middlewares

import (
	"app/internal/infrastructure/adapters"
	"app/internal/infrastructure/persistence/cache"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Two-Factor Authentication (2FA).
func Auth2FA() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")

		isValid, claims := adapters.Static.ValidateJWT(tokenString)
		if isValid == false {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Token inválido")
			return
		}

		userInSession, err := cache.Session.GetUserSession(claims.UserID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Não há usuário em sessão")
			return
		}

		if adapters.RemoveBearerPrefix(tokenString) != userInSession.JWT {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Token distinto")
			return
		}

		ctx.Set("user_session", userInSession)

		ctx.Next()
	}
}
