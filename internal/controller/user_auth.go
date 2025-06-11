package controller

import (
	"app/internal/dto"
	"app/internal/mytypes"
	"app/internal/userauth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (it *LayerController) LoginUser(ctx *gin.Context) {
	request, err := MapReqJSON[dto.Login](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	stringJWT, err := it.useCase.UserLogin(*request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, mytypes.ErrorRes{
			Status: http.StatusUnauthorized,
			Error:  err,
		})
		return
	}

	ctx.String(http.StatusOK, stringJWT)
}

func (it *LayerController) LogoutUser(ctx *gin.Context) {
	userInfo, err := userauth.GetInfoSession(ctx, "user_session")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = it.useCase.UserLogout(userInfo.Id)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.String(http.StatusOK, "Logout Ok")
}

func (it *LayerController) RegisterUser(ctx *gin.Context) {
	request, err := MapReqJSON[dto.UserRegisterReq](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = it.useCase.UserRegister(*request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, mytypes.ErrorRes{
			Status: http.StatusUnauthorized,
			Error:  err,
		})
		return
	}

	ctx.String(http.StatusCreated, "Usuário Criado")
}

func (it *LayerController) SendRefreshPassword(ctx *gin.Context) {
	EmailParam := ctx.Param("email")

	err := it.useCase.SendRefreshForEmail(EmailParam)
	if err != nil {
		ctx.String(http.StatusUnauthorized, "Erro de validação: "+err.Error())
		return
	}

	ctx.String(http.StatusCreated, "Refresh Password enviado")
}

func (it *LayerController) RefreshPassword(ctx *gin.Context) {
	stringJWT := ctx.Query("jwt")
	fmt.Println("s", stringJWT)
}
