// trata os dados das resquest/response
package controller

import (
	"app/internal/dto"
	"app/internal/userauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (it *UserController) CreateUser(ctx *gin.Context) {
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

func (it *UserController) GetUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	userInfo, err := GetInfoSession[userauth.UserSession](ctx, "user_session")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	if userInfo.Role != "admin" && paramID != userInfo.Id {
		ctx.JSON(403, "Acesso não autorizado")
		return
	}

	response, err := it.useCase.GetUser(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (it *UserController) GetAllUsers(ctx *gin.Context) {
	userInfo, err := GetInfoSession[userauth.UserSession](ctx, "user_session")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	if userInfo.Role != "admin" {
		ctx.JSON(403, "Acesso não autorizado")
		return
	}

	response, err := it.useCase.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (it *UserController) UpdateUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	request, err := MapReqJSON[dto.UserUpdateReq](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	request.ID = paramID

	userInfo, err := GetInfoSession[userauth.UserSession](ctx, "user_session")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	if userInfo.Role != "admin" && paramID != userInfo.Id {
		ctx.JSON(403, "Acesso não autorizado")
		return
	}

	err = it.useCase.UpdateUser(*request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.String(http.StatusOK, "Usuário atualizado com sucesso")
}

func (it *UserController) DeleteUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	userInfo, err := GetInfoSession[userauth.UserSession](ctx, "user_session")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	if userInfo.Role != "admin" && paramID != userInfo.Id {
		ctx.JSON(403, "Acesso não autorizado")
		return
	}

	err = it.useCase.DeleteUser(paramID)
	it.useCase.UserLogout(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.String(http.StatusNoContent, "Usuário deletado com sucesso")
}
