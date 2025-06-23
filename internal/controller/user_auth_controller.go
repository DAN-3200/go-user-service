package controller

import (
	"app/internal/dto"
	"app/internal/mytypes"
	"app/internal/userauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (it *LayerController) LoginUser(ctx *gin.Context) {
	request, err := MapReqJSON[dto.Login](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	stringJWT, err := it.useCase.LoginUser(*request)
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

	err = it.useCase.LogoutUser(userInfo.Id)
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

	err = it.useCase.RegisterUser(*request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, mytypes.ErrorRes{
			Status: http.StatusUnauthorized,
			Error:  err,
		})
		return
	}

	ctx.String(http.StatusCreated, "Usuário Criado")
}

func (it *LayerController) SendRefreshForEmail(ctx *gin.Context) {
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

	// validar se expirou
	isValid, claims := userauth.ValidateJWT(stringJWT)
	if !isValid {
		ctx.String(http.StatusUnauthorized, "JWT inválido")
		return
	}

	request, err := MapReqJSON[dto.RefreshPassword](ctx)
	if err != nil {
		ctx.String(http.StatusUnauthorized, err.Error())
		return
	}

	err = it.useCase.RefreshPassword(claims.UserID,*request)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.String(http.StatusOK, "Senha redefinida com sucesso!")
}
