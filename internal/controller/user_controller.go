// trata os dados das resquest/response
package controller

import (
	"app/internal/model"
	"app/internal/mytypes"
	"app/internal/usecase"
	"app/internal/userauth"
	"fmt"
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

	ctx.JSON(http.StatusOK, "Ok")
}

func (it *UserController) ReadUser(ctx *gin.Context) {
	// tratamento da body request
	var requestJWT = ctx.Request.Header.Get("Authorization")

	// Validação do JWT
	validate, claims := userauth.ValidateJWT(requestJWT)
	if validate == false {
		ctx.JSON(http.StatusUnauthorized, "JWT Não validado")
		return
	}

	// Verificação do Token em Sessão em relação do JWT fornecido
	userInSession, err := userauth.GetUserSession(claims.UserID)
	if err != nil {
		fmt.Println("Get User invalid:", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if userauth.RemoveBearerPrefix(requestJWT) != userInSession.JWT {
		ctx.JSON(http.StatusUnauthorized, "Token de autorização inválido")
		return
	}

	idParam, err := strconv.Atoi(claims.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Id inválido")
	}

	var response = it.useCase.UserRead(idParam)
	ctx.JSON(http.StatusOK, response)
}

func (it *UserController) ReadAllUser(ctx *gin.Context) {
	var requestJWT = ctx.Request.Header.Get("Authorization")

	// Validação do JWT
	validate, claims := userauth.ValidateJWT(requestJWT)
	if validate == false {
		ctx.JSON(http.StatusUnauthorized, "JWT Não validado")
		return
	}

	if claims.Role != "admin" {
		ctx.JSON(401, "Acesso não autorizado")
		return
	}

	// Verificação do Token em Sessão em relação do JWT fornecido
	userInSession, err := userauth.GetUserSession(claims.UserID)
	if err != nil {
		fmt.Println("Get User invalid:", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if userauth.RemoveBearerPrefix(requestJWT) != userInSession.JWT {
		ctx.JSON(http.StatusUnauthorized, "Token de autorização inválido")
		return
	}

	response, err := it.useCase.ReadAllUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
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

	var requestJWT = ctx.Request.Header.Get("Authorization")

	validate, claims := userauth.ValidateJWT(requestJWT)
	if validate == false {
		ctx.JSON(http.StatusUnauthorized, "JWT Não validado")
		return
	}

	userInSession, err := userauth.GetUserSession(claims.UserID)
	if err != nil {
		fmt.Println("Get User invalid:", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if claims.Role != "admin" && request.Id != userInSession.Id {
		ctx.JSON(401, "Acesso não autorizado")
		return
	}

	if userauth.RemoveBearerPrefix(requestJWT) != userInSession.JWT {
		ctx.JSON(http.StatusUnauthorized, "Token de autorização inválido")
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

	var requestJWT = ctx.Request.Header.Get("Authorization")

	validate, claims := userauth.ValidateJWT(requestJWT)
	if validate == false {
		ctx.JSON(http.StatusUnauthorized, "JWT Não validado")
		return
	}

	userInSession, err := userauth.GetUserSession(claims.UserID)
	if err != nil {
		fmt.Println("Get User invalid:", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if claims.Role != "admin" && pID != userInSession.Id {
		ctx.JSON(401, "Acesso não autorizado")
		return
	}

	if userauth.RemoveBearerPrefix(requestJWT) != userInSession.JWT {
		ctx.JSON(http.StatusUnauthorized, "Token de autorização inválido")
		return
	}

	err = it.useCase.UserDelete(pID)
	it.useCase.UserLogout(strconv.Itoa(pID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "User Deleted")
}
