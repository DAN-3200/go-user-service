// trata os dados das resquest/response
package controller

import (
	"app/internal/model"
	"app/internal/mytypes"
	"app/internal/usecase"
	"app/internal/userauth"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	useCase usecase.UserUseCase
}

func NewUserController(useCase usecase.UserUseCase) *UserController {
	return &UserController{useCase}
}

// -- Methods

func (it *UserController) CreateUser(ctx *gin.Context) {
	var request model.User
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, "Erro na leitura da requisição")
		return
	}
	err := request.Validate()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, mytypes.DetailsError{
			HttpStatus: http.StatusBadRequest,
			Error:      err.Error(),
		})
		return
	}

	err = it.useCase.UserCreate(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, "Ok")
}

func (it *UserController) ReadUser(ctx *gin.Context) {
	paramID := ctx.Param("id")
	pID, _ := strconv.Atoi(paramID)

	var response = it.useCase.UserRead(pID)
	ctx.JSON(http.StatusOK, response)
}

func (it *UserController) ReadAllUser(ctx *gin.Context) {
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

	response, err := it.useCase.ReadAllUser()
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (it *UserController) UpdateUser(ctx *gin.Context) {
	var request model.User
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, mytypes.DetailsError{
			HttpStatus: http.StatusBadRequest,
			Error:      "Erro na leitura da requisição",
		})
		return
	}

	err := request.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, mytypes.DetailsError{
			HttpStatus: http.StatusBadRequest,
			Error:      err.Error(),
		})
		return
	}

	value, _ := ctx.Get("user_session")
	userInfo, ok := value.(*userauth.UserSession)
	if !ok {
		log.Println("Falha na conversão para userauth.UserSession")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Erro interno no servidor")
		return
	}

	if userInfo.Role != "admin" && request.Id != userInfo.Id {
		ctx.JSON(403, "Acesso não autorizado")
		return
	}

	err = it.useCase.UserUpdate(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Ok")
}

func (it *UserController) DeleteUser(ctx *gin.Context) {
	paramID := ctx.Param("id")
	pID, _ := strconv.Atoi(paramID)

	value, _ := ctx.Get("user_session")
	userInfo, ok := value.(*userauth.UserSession)
	if !ok {
		log.Println("Falha na conversão para userauth.UserSession")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Erro interno no servidor")
		return
	}

	if userInfo.Role != "admin" && pID != userInfo.Id {
		ctx.JSON(403, "Acesso não autorizado")
		return
	}

	err := it.useCase.UserDelete(pID)
	it.useCase.UserLogout(strconv.Itoa(pID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, "User Deleted")
}
