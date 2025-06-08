package usecase

import "app/internal/contracts"

type LayerUseCase struct {
	Repo contracts.UserRepoSQL
}

func Init(repo contracts.UserRepoSQL) *LayerUseCase {
	return &LayerUseCase{repo}
}

// ------------------------------------------------------------------------
