package controller

import (
	"app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type LayerController struct {
	useCase *usecase.LayerUseCase
}

func Init(useCase *usecase.LayerUseCase) *LayerController {
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
