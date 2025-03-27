// trata os dados das resquest/response
package controller

import (
	"app/model"
	"app/useCase"
	"app/userAuth"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	useCase useCase.UserUseCase
}

func NewUserController(useCase useCase.UserUseCase) *UserController {
	return &UserController{useCase}
}

// -- Methods

func (it *UserController) ReadUser(ctx *gin.Context) {
	// tratamento da body request
	var requestJWT = ctx.Request.Header.Get("Authorization")

	// Validação do JWT
	validate, claims := userAuth.ValidateJWT(requestJWT)
	if validate == false {
		ctx.JSON(http.StatusUnauthorized, "JWT Não validado")
		return
	}

	// Verificação do Token em Sessão em relação do JWT fornecido
	userInSession, err := userAuth.GetUserSession(claims.UserID)
	if err != nil {
		fmt.Println("Get User invalid:", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if userAuth.RemoveBearerPrefix(requestJWT) != userInSession.JWT {
		ctx.JSON(http.StatusUnauthorized, "Token de autorização inválido")
		return
	}

	idParam, err := strconv.Atoi(claims.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Id inválido")
	}
	var response = it.useCase.UserRead(idParam)
	ctx.JSON(http.StatusOK, response)
}

func (it *UserController) ReadAllUser(ctx *gin.Context) {
	var requestJWT = ctx.Request.Header.Get("Authorization")

	// Validação do JWT
	validate, claims := userAuth.ValidateJWT(requestJWT)
	if validate == false {
		ctx.JSON(http.StatusUnauthorized, "JWT Não validado")
		return
	}

	if claims.Role != "admin" {
		ctx.JSON(401, "Acesso não autorizado")
		return
	}

	// Verificação do Token em Sessão em relação do JWT fornecido
	userInSession, err := userAuth.GetUserSession(claims.UserID)
	if err != nil {
		fmt.Println("Get User invalid:", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if userAuth.RemoveBearerPrefix(requestJWT) != userInSession.JWT {
		ctx.JSON(http.StatusUnauthorized, "Token de autorização inválido")
		return
	}

	response, err := it.useCase.ReadAllUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (it *UserController) CreateUser(ctx *gin.Context) {
	var request model.User
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, "Erro na leitura da requisição")
		return
	}

	var err = it.useCase.UserCreate(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Ok")
}

// tem erro aqui
func (it *UserController) UpdateUser(ctx *gin.Context) {
	var request model.User
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, "Erro na leitura da requisição")
		return
	}

	var err = it.useCase.UserUpdate(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Ok")
}

func (it *UserController) DeleteUser(ctx *gin.Context) {
	var request struct{ Id int }
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, "Erro na leitura da requisição")
		return
	}

	var err = it.useCase.UserDelete(request.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Ok")

}
