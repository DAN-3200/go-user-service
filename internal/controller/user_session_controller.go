package controller

import (
	"app/internal/model"
	"app/internal/mytypes"
	"app/internal/userauth"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (it *UserController) LoginUser(ctx *gin.Context) {
	var request model.LoginFields
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusInternalServerError, "Erro na leitura da requisição")
		return
	}

	errList, err := request.ValidateFields()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, mytypes.DetailsError{
			HttpStatus: http.StatusBadRequest,
			Error:      errList,
		})
		return
	}

	stringJWT, err := it.useCase.UserLogin(request.Email, request.Password)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, any(stringJWT))
}

func (it *UserController) LogoutUser(ctx *gin.Context) {
	value, _ := ctx.Get("user_session")
	userInfo, ok := value.(*userauth.UserSession)
	if !ok {
		log.Println("Falha na conversão para userauth.UserSession")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Erro interno no servidor")
		return
	}

	// Remoção do User da sessão
	err := it.useCase.UserLogout(strconv.Itoa(userInfo.Id))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, "Logout Ok")
}
