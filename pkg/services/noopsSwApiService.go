package services

import (
	"context"

	"github.com/peterhellberg/swapi"
)

type (
	noOpsSwApiService struct {
		sm ServiceManager
	}
)

func NewNoOpsSwService(ctx context.Context) SwApiService {
	return &noOpsSwApiService{}
}

func (n *noOpsSwApiService) Start(ctx context.Context) error {
	return nil
}

func (n *noOpsSwApiService) Close(ctx context.Context) error {
	return nil
}

func (n *noOpsSwApiService) Healthy(ctx context.Context) error {
	return nil
}

func (n *noOpsSwApiService) WithServiceManager(sm ServiceManager) SwApiService {
	n.sm = sm
	return n
}

func (n *noOpsSwApiService) ServiceManager() ServiceManager {
	return n.sm
}

func (n *noOpsSwApiService) GetPlanetById(id int) (*swapi.Planet, error) {
	return &swapi.Planet{
		Name: "DummyPlanet",
	}, nil
}
