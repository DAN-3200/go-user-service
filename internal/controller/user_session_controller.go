package controller

import (
	"app/internal/model"
	"app/internal/mytypes"
	"app/internal/userauth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (it *UserController) LoginUser(ctx *gin.Context) {
	var request model.LoginFields
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, "Erro na leitura da requisição")
		return
	}

	errList, err := request.ValidateFields()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, mytypes.DetailsError{
			HttpStatus: http.StatusBadRequest,
			Error:      errList,
		})
		return
	}

	stringJWT, err := it.useCase.UserLogin(request.Email, request.Password)
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
