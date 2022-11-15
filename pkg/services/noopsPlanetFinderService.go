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

func NewNoOpsPlanetFinderService(ctx context.Context) PlanetFinderService {
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
