package services

import (
	"context"

	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/peterhellberg/swapi"
)

type (
	noOpsSwApiService struct {
		sm ServiceManager
	}
)

func NewNoOpsSwService() SwApiService {
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

func (n *noOpsSwApiService) GetPlanetById(ctx context.Context, id int) (*models.Planet, error) {
	return &models.Planet{
		Name: "DummyPlanet",
	}, nil
}

func (n *noOpsSwApiService) SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error) {
	return EmptyPlanetSlice, nil
}

func (n *noOpsSwApiService) ToPersistentPlanet(ctx context.Context, p *swapi.Planet, id int, expand bool) (*models.Planet, error) {
	return &models.Planet{
		Name: "DummyPlanet",
	}, nil
}
func (n *noOpsSwApiService) ToPersistentFilm(ctx context.Context, f *swapi.Film, id int, expand bool) (*models.Film, error) {
	return &models.Film{
		Title: "DummyFilm",
	}, nil
}

func (n *noOpsSwApiService) PutOnline() SwApiService {
	return n
}
func (n *noOpsSwApiService) PutOffline() SwApiService {
	return n
}

func (n *noOpsSwApiService) IsOnline() bool {
	return false
}
