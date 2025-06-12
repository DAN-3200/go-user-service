package usecase

import (
	"app/internal/contracts"
)

type LayerUseCase struct {
	Repo  contracts.UserRepoSQL
	Drive contracts.Drive
}

func Init(repo contracts.UserRepoSQL, drive contracts.Drive) *LayerUseCase {
	return &LayerUseCase{repo, drive}
}
