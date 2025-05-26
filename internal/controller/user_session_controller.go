package controller

import (
	"app/internal/model"
	"app/internal/mytypes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (it *UserController) LoginUser(ctx *gin.Context) {
	var request model.LoginFields
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, "Erro na leitura da requisição")
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
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, stringJWT)
}

func (it *UserController) LogoutUser(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	
	// Remoção do User da sessão
	err := it.useCase.UserLogout(userID.(string))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Logout Ok")
}
