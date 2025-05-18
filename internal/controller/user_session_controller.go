package controller

import (
	"app/internal/userauth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (it *UserController) LoginUser(ctx *gin.Context) {
	var request struct {
		UserEmail    string `json:"email"`
		UserPassword string `json:"password"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, "Erro na leitura da requisição")
		return
	}

	var stringJWT, err = it.useCase.UserLogin(request.UserEmail, request.UserPassword)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, stringJWT)
}

func (it *UserController) LogoutUser(ctx *gin.Context) {
	// tratamento da body request
	var requestJWT = ctx.Request.Header.Get("Authorization")

	// Validação do JWT
	validate, claims := userauth.ValidateJWT(requestJWT)
	if validate == false {
		ctx.JSON(http.StatusUnauthorized, "JWT Não validado")
		return
	}

	// Verificação do Token em Sessão em relação do JWT fornecido
	userInSession, err := userauth.GetUserSession(claims.UserID)
	if err != nil {
		fmt.Println("Get User invalid:", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if userauth.RemoveBearerPrefix(requestJWT) != userInSession.JWT {
		ctx.JSON(http.StatusUnauthorized, "Token de autorização inválido")
		return
	}

	// Remoção do User da sessão
	err = it.useCase.UserLogout(claims.UserID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Logout Ok")
}
