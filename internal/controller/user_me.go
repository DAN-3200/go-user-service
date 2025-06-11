package controller

import (
	"app/internal/dto"
	"app/internal/userauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (it *LayerController) GetMyInfo(ctx *gin.Context) {
	userInfo, err := userauth.GetInfoSession(ctx, "user_session")
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

func (it *LayerController) EditMyInfo(ctx *gin.Context) {
	request, err := MapReqJSON[dto.EditMeReq](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	userInfo, err := userauth.GetInfoSession(ctx, "user_session")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	request.ID = userInfo.Id

	// err = it.useCase.EditMyInfo(*request)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	ctx.JSON(http.StatusOK, "informações pessoais editada")
}
