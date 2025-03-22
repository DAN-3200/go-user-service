// trata os dados das resquest/response
package controller

import (
	"app/model"
	"app/useCase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	useCase useCase.UserUseCase
}

func NewUserController(useCase useCase.UserUseCase) *UserController {
	return &UserController{useCase}
}

// -- Methods

func (it *UserController) UserRead(ctx *gin.Context) {
	var response = it.useCase.UserRead()
	ctx.JSON(http.StatusOK, response)
}
func (it *UserController) UserCreate(ctx *gin.Context) {
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
