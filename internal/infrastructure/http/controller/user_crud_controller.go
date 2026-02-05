// trata os dados das resquest/response
package controller

import (
	"app/internal/application/dto"
	"app/internal/application/usecase"
	"app/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


type UserManagerController struct {
	useCase *usecase.UserManagerUC
}

func InitUserManager(usecase *usecase.UserManagerUC) *UserManagerController {
	return &UserManagerController{usecase}
}


func (it *UserManagerController) CreateUser(ctx *gin.Context) {
	request, err := utils.MapReqJSON[dto.UserRegisterReq](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = it.useCase.RegisterUser(*request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.String(http.StatusCreated, "Usuário criado com sucesso")
}

func (it *UserManagerController) GetUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	response, err := it.useCase.GetUser(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (it *UserManagerController) GetUserList(ctx *gin.Context) {
	response, err := it.useCase.GetUserList()
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (it *UserManagerController) EditUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	request, err := utils.MapReqJSON[dto.EditUserReq](ctx)
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

func (it *UserManagerController) DeleteUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	err := it.useCase.DeleteUser(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.String(http.StatusNoContent, "Usuário deletado com sucesso")
}
