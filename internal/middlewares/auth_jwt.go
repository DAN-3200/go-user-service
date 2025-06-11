package middlewares

import (
	"app/internal/userauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Two-Factor Authentication (2FA).
func Auth2FA() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")

		isValid, claims := userauth.ValidateJWT(tokenString)
		if isValid == false {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "JWT inválido")
			return
		}

		userInSession, err := userauth.GetUserSession(claims.UserID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Não há usuário em sessão")
			return
		}

		if userauth.RemoveBearerPrefix(tokenString) != userInSession.JWT {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Distinct token")
			return
		}

		ctx.Set("user_session", userInSession)

		ctx.Next()
	}
}
