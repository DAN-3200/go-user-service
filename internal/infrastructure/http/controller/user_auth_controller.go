package controller

import (
	"app/internal/application/dto"
	"app/internal/application/usecase"
	"app/internal/infrastructure/adapters"
	"app/internal/infrastructure/persistence/cache"
	"app/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAuthController struct {
	useCase *usecase.UserAuthUC
}

func InitUserAuth(usecase *usecase.UserAuthUC) *UserAuthController {
	return &UserAuthController{usecase}
}

func (it *UserAuthController) LoginUser(ctx *gin.Context) {
	request, err := utils.MapReqJSON[dto.Login](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	stringJWT, err := it.useCase.LoginUser(*request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.String(http.StatusOK, stringJWT)
}

func (it *UserAuthController) LogoutUser(ctx *gin.Context) {
	userInfo, err := cache.Session.GetInfoSession(ctx, "user_session")
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

func (it *UserAuthController) SendRefreshForEmail(ctx *gin.Context) {
	EmailParam := ctx.Param("email")

	err := it.useCase.SendRefreshForEmail(EmailParam)
	if err != nil {
		ctx.String(http.StatusUnauthorized, "Erro de validação: "+err.Error())
		return
	}

	ctx.String(http.StatusCreated, "Refresh Password enviado")
}

func (it *UserAuthController) RefreshPassword(ctx *gin.Context) {
	stringJWT := ctx.Query("jwt")

	// validar se expirou
	isValid, claims := adapters.Static.ValidateJWT(stringJWT)
	if !isValid {
		ctx.String(http.StatusUnauthorized, "JWT inválido")
		return
	}

	request, err := utils.MapReqJSON[dto.RefreshPassword](ctx)
	if err != nil {
		ctx.String(http.StatusUnauthorized, err.Error())
		return
	}

	err = it.useCase.RefreshPassword(claims.UserID, *request)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.String(http.StatusOK, "Senha redefinida com sucesso!")
}
