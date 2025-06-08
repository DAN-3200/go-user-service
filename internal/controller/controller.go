package controller

import (
	"app/internal/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

type LayerController struct {
	useCase usecase.LayerUseCase
}

func Init(useCase usecase.LayerUseCase) *LayerController {
	return &LayerController{useCase}
}

// ------------------------------------------------------------------------

func MapReqJSON[T any](ctx *gin.Context) (*T, error) {
	var request T
	if err := ctx.BindJSON(&request); err != nil {
		return &request, err
	}
	return &request, nil
}

func GetInfoSession[T any](ctx *gin.Context, key string) (*T, error) {
	var value_format *T

	value_base, ok := ctx.Get(key)
	if !ok {
		return value_format, fmt.Errorf("Key não existe")
	}

	value_format, ok = value_base.(*T)
	if !ok {
		return value_format, fmt.Errorf("Erro ao formatar a informação")
	}
	return value_format, nil
}
