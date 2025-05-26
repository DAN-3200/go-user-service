package middlewares

import (
	"app/internal/userauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthJWT(SaveSession bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// tratamento da body request
		var requestJWT = ctx.Request.Header.Get("Authorization")

		// Validação do JWT
		validate, claims := userauth.ValidateJWT(requestJWT)
		if validate == false {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "JWT invalid")

			return
		}

		// Verificação do Token em Sessão em relação do JWT fornecido
		userInSession, err := userauth.GetUserSession(claims.UserID)
		if err != nil {
			// fmt.Println("No user in session:", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "No user in session")

			return
		}

		if userauth.RemoveBearerPrefix(requestJWT) != userInSession.JWT {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Distinct token")

			return
		}

		ctx.Set("user_session", userInSession)

		ctx.Next()
	}
}
