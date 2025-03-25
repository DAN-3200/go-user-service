package controller

import (
	"app/userAuth"
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

	ctx.JSON(http.StatusOK, struct{ JWT string }{stringJWT})
}

func (it *UserController) LogoutUser(ctx *gin.Context) {
	// tratamento da body request
	var request struct {
		JWT string `json:"JWT"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, "Erro na leitura da requisição")
		return
	}

	// Validação do JWT
	validate, claims := userAuth.ValidateJWT(request.JWT)
	if validate == false {
		ctx.JSON(http.StatusUnauthorized, "JWT Não validado")
		return
	}

	// Verificação do Token em Sessão em relação do JWT fornecido
	userID := claims["userId"].(string)
	userInSession, err := userAuth.GetUserSession(userID)
	if err != nil {
		fmt.Println("Get User invalid:", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if request.JWT != userInSession.JWT {
		ctx.JSON(http.StatusUnauthorized, "Token de autorização inválido")
		return
	}

	// Remoção do User da sessão
	err = it.useCase.UserLogout(userID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Logout Ok")
}
