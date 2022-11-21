package services

import (
	"context"
	"errors"

	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/peterhellberg/swapi"
)

type (
	noOpsSwApiService struct {
		sm     ServiceManager
		online bool
	}
)

func NewNoOpsSwService() SwApiService {
	return &noOpsSwApiService{online: true}
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
	if id == 1 {
		return &models.Planet{
			Id:      1,
			Name:    "Terra",
			Climate: "Tropical",
			Terrain: "continental",
			URL:     "https://something.com/api/planet/1/",
		}, nil
	} else if id == 0 {
		return nil, errors.New("dummy Error")
	}
	return nil, nil
}

func (n *noOpsSwApiService) SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error) {
	if name != "error" {
		return []*models.Planet{
			{
				Id:      1,
				Name:    "Terra",
				Climate: "Tropical",
				Terrain: "continental",
				URL:     "https://something.com/api/planet/1/",
			},
			{
				Id:      2,
				Name:    "Marte",
				Climate: "Hazard",
				Terrain: "Blob",
				URL:     "https://something.com/api/planet/2/",
			},
		}, nil
	}
	return EmptyPlanetSlice, errors.New("some Error")
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
	n.online = true
	return n
}
func (n *noOpsSwApiService) PutOffline() SwApiService {
	n.online = false
	return n
}

func (n *noOpsSwApiService) IsOnline() bool {
	return n.online
}
