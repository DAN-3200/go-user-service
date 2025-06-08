package controller

import (
	"app/internal/dto"
	"app/internal/mytypes"
	"app/internal/userauth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (it *UserController) LoginUser(ctx *gin.Context) {
	request, err := MapReqJSON[dto.Login](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = request.ValidateFields()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, mytypes.ErrorRes{
			Status: http.StatusBadRequest,
			Error:  err,
		})
		return
	}

	stringJWT, err := it.useCase.UserLogin(*request)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, mytypes.ErrorRes{
			Status: http.StatusUnauthorized,
			Error:  err,
		})
		return
	}

	ctx.String(http.StatusOK, stringJWT)
}

func (it *UserController) LogoutUser(ctx *gin.Context) {
	value, _ := ctx.Get("user_session")
	userInfo, ok := value.(*userauth.UserSession)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Falha na conversão para userauth.UserSession")
		return
	}

	err := it.useCase.UserLogout(userInfo.Id)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, "Logout Ok")
}

func (it *UserController) RegisterUser(ctx *gin.Context) {
	request, err := MapReqJSON[dto.UserRegisterRes](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = it.useCase.UserRegister(*request)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, mytypes.ErrorRes{
			Status: http.StatusUnauthorized,
			Error:  err,
		})
		return
	}

	ctx.String(http.StatusCreated, "Usuário Criado")
}
