package controller

import (
	"app/internal/domain/usecase"

	"github.com/gin-gonic/gin"
)

type LayerController struct {
	useCase *usecase.LayerUseCase
}

func Init(usecase *usecase.LayerUseCase) *LayerController {
	return &LayerController{usecase}
}

// ------------------------------------------------------------------------

func MapReqJSON[T any](ctx *gin.Context) (*T, error) {
	var request T
	if err := ctx.BindJSON(&request); err != nil {
		return &request, err
	}
	return &request, nil
}
