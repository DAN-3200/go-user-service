// trata os dados das resquest/response
package controller

import (
	"app/internal/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (it *LayerController) CreateUser(ctx *gin.Context) {
	request, err := MapReqJSON[dto.UserReq](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = it.useCase.CreateUser(*request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.String(http.StatusCreated, "Usuário criado com sucesso")
}

func (it *LayerController) GetUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	response, err := it.useCase.GetUser(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (it *LayerController) GetUserList(ctx *gin.Context) {
	response, err := it.useCase.GetUserList()
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (it *LayerController) EditUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	request, err := MapReqJSON[dto.EditUserReq](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = it.useCase.EditUser(paramID, *request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.String(http.StatusOK, "Usuário atualizado com sucesso")
}

func (it *LayerController) DeleteUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	err := it.useCase.DeleteUser(paramID)
	it.useCase.LogoutUser(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.String(http.StatusNoContent, "Usuário deletado com sucesso")
}
