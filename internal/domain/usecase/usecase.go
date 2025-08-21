package usecase

import "app/internal/domain/ports"

type LayerUseCase struct {
	Repo    ports.RepositoryPorts
	Service ports.ServicePorts
}

func Init(repo ports.RepositoryPorts, service ports.ServicePorts) *LayerUseCase {
	return &LayerUseCase{repo, service}
}
