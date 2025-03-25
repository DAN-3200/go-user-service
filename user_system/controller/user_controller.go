// trata os dados das resquest/response
package controller

import (
	"app/model"
	"app/useCase"
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
	request := ctx.Param("id")
	if request == "" {
		ctx.JSON(http.StatusBadRequest, "Id não fornecido")
	}

	idParam, err := strconv.Atoi(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Id inválido")
	}
	var response = it.useCase.UserRead(idParam)
	ctx.JSON(http.StatusOK, response)
}

func (it *UserController) ReadAllUser(ctx *gin.Context) {
	var response, err = it.useCase.ReadAllUser()
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