package controller

import (
	"app/internal/userauth"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (it *UserController) MyInfo(ctx *gin.Context) {
	value, _ := ctx.Get("user_session")
	userInSession, ok := value.(*userauth.UserSession)
	if !ok {
		log.Println("Falha na conversão para userauth.UserSession")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Erro interno no servidor")
		return
	}

	response, err := it.useCase.MyInfo(userInSession.Id)
	if err != nil {
		fmt.Println("Error:", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
