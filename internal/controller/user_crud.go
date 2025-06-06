// trata os dados das resquest/response
package controller

import (
	"app/internal/dto"
	"app/internal/mytypes"
	"app/internal/usecase"
	"app/internal/userauth"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	useCase usecase.UserUseCase
}

func NewUserController(useCase usecase.UserUseCase) *UserController {
	return &UserController{useCase}
}

// ------------------------------------------------------------------------

func (it *UserController) CreateUser(ctx *gin.Context) {
	var request dto.UserReq
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, "Erro na leitura da requisição")
		return
	}

	err := request.ValidateFields()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = it.useCase.CreateUser(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, "Usuário criado com sucesso")
}

func (it *UserController) GetUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	response, err := it.useCase.GetUser(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	value, _ := ctx.Get("user_session")
	userInfo, ok := value.(*userauth.UserSession)
	if !ok {
		log.Println("Falha na conversão para userauth.UserSession")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Erro interno no servidor")
		return
	}

	if userInfo.Role != "admin" && paramID != userInfo.Id {
		ctx.JSON(403, "Acesso não autorizado")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (it *UserController) GetAllUsers(ctx *gin.Context) {
	value, _ := ctx.Get("user_session")
	userInfo, ok := value.(*userauth.UserSession)
	if !ok {
		log.Println("Falha na conversão para userauth.UserSession")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Erro interno no servidor")
		return
	}

	if userInfo.Role != "admin" {
		ctx.JSON(403, "Acesso não autorizado")
		return
	}

	response, err := it.useCase.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (it *UserController) UpdateUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	var request dto.UserUpdateReq
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, mytypes.ErrorRes{
			Status: http.StatusBadRequest,
			Error:  fmt.Errorf("Erro na leitura da requisição"),
		})
		return
	}
	request.ID = paramID

	value, _ := ctx.Get("user_session")
	userInfo, ok := value.(*userauth.UserSession)
	if !ok {
		log.Println("Falha na conversão para userauth.UserSession")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Erro interno no servidor")
		return
	}

	if userInfo.Role != "admin" && paramID != userInfo.Id {
		ctx.JSON(403, "Acesso não autorizado")
		return
	}

	err := it.useCase.UpdateUser(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Usuário atualizado com sucesso")
}

func (it *UserController) DeleteUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	value, _ := ctx.Get("user_session")
	userInfo, ok := value.(*userauth.UserSession)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Erro interno no servidor")
		return
	}

	if userInfo.Role != "admin" && paramID != userInfo.Id {
		ctx.JSON(403, "Acesso não autorizado")
		return
	}

	err := it.useCase.DeleteUser(paramID)
	it.useCase.UserLogout(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, "Usuário deletado com sucesso")
}
