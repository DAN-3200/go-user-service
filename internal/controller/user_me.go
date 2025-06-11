package controller

import (
	"app/internal/userauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (it *LayerController) MyInfo(ctx *gin.Context) {
	userInfo, err := userauth.GetInfoSession(ctx, "user_session")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	response, err := it.useCase.MyInfo(userInfo.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
