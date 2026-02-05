package controller

import (
	"app/internal/application/dto"
	"app/internal/application/usecase"
	"app/internal/infrastructure/persistence/cache"
	"app/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfoController struct {
	useCase *usecase.UserInfoUC
}

func InitUserInfo(usecase *usecase.UserInfoUC) *UserInfoController {
	return &UserInfoController{usecase}
}

func (it *UserInfoController) GetMyInfo(ctx *gin.Context) {
	userInfo, err := cache.Session.GetInfoSession(ctx, "user_session")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	response, err := it.useCase.GetMyInfo(userInfo.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (it *UserInfoController) EditMyInfo(ctx *gin.Context) {
	request, err := utils.MapReqJSON[dto.EditMeReq](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	userInfo, err := cache.Session.GetInfoSession(ctx, "user_session")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = it.useCase.EditMyInfo(userInfo.Id, *request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.String(http.StatusOK, "informações pessoais editada")
}
