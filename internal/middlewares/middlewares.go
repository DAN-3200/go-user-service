package middlewares

import (
	"app/internal/userauth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// tratamento da body request
		var requestJWT = ctx.Request.Header.Get("Authorization")

		// Validação do JWT
		validate, claims := userauth.ValidateJWT(requestJWT)
		if validate == false {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "JWT Não validado")

			return
		}

		// Verificação do Token em Sessão em relação do JWT fornecido
		userInSession, err := userauth.GetUserSession(claims.UserID)
		if err != nil {
			fmt.Println("Get User invalid:", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)

			return
		}

		if userauth.RemoveBearerPrefix(requestJWT) != userInSession.JWT {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Token de autorização inválido")

			return
		}

		ctx.Next()
	}
}

func AuthForLogout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// tratamento da body request
		var requestJWT = ctx.Request.Header.Get("Authorization")

		// Validação do JWT
		validate, claims := userauth.ValidateJWT(requestJWT)
		if validate == false {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "JWT Não validado")
			return
		}

		// Verificação do Token em Sessão em relação do JWT fornecido
		userInSession, err := userauth.GetUserSession(claims.UserID)
		if err != nil {
			fmt.Println("Get User invalid:", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		if userauth.RemoveBearerPrefix(requestJWT) != userInSession.JWT {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Token de autorização inválido")
			return
		}
		
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
