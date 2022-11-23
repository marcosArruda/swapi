package services

import (
	"context"

	"github.com/marcosArruda/swapi/pkg/models"
)

type (
	noOpsPlanetFinderService struct {
		sm ServiceManager
	}
)

func NewNoOpsPlanetFinderService() PlanetFinderService {
	return &noOpsPlanetFinderService{}
}

func (n *noOpsPlanetFinderService) Start(ctx context.Context) error {
	return nil
}

func (n *noOpsPlanetFinderService) Close(ctx context.Context) error {
	return nil
}

func (n *noOpsPlanetFinderService) Healthy(ctx context.Context) error {
	return nil
}

func (n *noOpsPlanetFinderService) WithServiceManager(sm ServiceManager) PlanetFinderService {
	n.sm = sm
	return n
}

func (n *noOpsPlanetFinderService) ServiceManager() ServiceManager {
	return n.sm
}

func (n *noOpsPlanetFinderService) GetPlanetById(ctx context.Context, id int) (*models.Planet, error) {
	if id == 1 {
		return &models.Planet{
			Id:       1,
			Name:     "Terra",
			Climate:  "tropical",
			Terrain:  "terra",
			FilmURLs: []string{"https://something.com/api/films/1/"},
			URL:      "https://something.com/api/planets/1/",
			Films: []*models.Film{
				{
					Id:         1,
					Title:      "Filme da Terra",
					EpisodeID:  1,
					Director:   "Único",
					Created:    "800 quintilhões de anos atras",
					PlanetURLs: []string{"https://something.com/api/planets/1/"},
					Planets:    EmptyPlanetSlice,
					URL:        "https://something.com/api/films/1/",
				},
			},
		}, nil
	}
	return nil, nil
}

func (n *noOpsPlanetFinderService) SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error) {
	return EmptyPlanetSlice, nil
}

func (n *noOpsPlanetFinderService) ListAllPlanets(ctx context.Context) ([]*models.Planet, error) {
	return EmptyPlanetSlice, nil
}

func (n *noOpsPlanetFinderService) RemovePlanetById(ctx context.Context, id int) error {
	return nil
}

func (n *noOpsPlanetFinderService) RemovePlanetByExactName(ctx context.Context, exactName string) error {
	return nil
}
