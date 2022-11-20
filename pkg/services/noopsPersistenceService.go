package services

import (
	"context"

	"github.com/marcosArruda/swapi/pkg/models"
)

type (
	noOpsPersistenceService struct {
		sm ServiceManager
	}
)

func NewNoOpsPersistenceService() PersistenceService {
	return &noOpsPersistenceService{}
}

func (n *noOpsPersistenceService) Start(ctx context.Context) error {
	return nil
}

func (n *noOpsPersistenceService) Close(ctx context.Context) error {
	return nil
}

func (n *noOpsPersistenceService) Healthy(ctx context.Context) error {
	return nil
}

func (n *noOpsPersistenceService) WithServiceManager(sm ServiceManager) PersistenceService {
	n.sm = sm
	return n
}

func (n *noOpsPersistenceService) ServiceManager() ServiceManager {
	return n.sm
}

func (n *noOpsPersistenceService) UpsertPlanet(ctx context.Context, p *models.Planet) error {
	return nil
}

func (n *noOpsPersistenceService) GetPlanetById(ctx context.Context, id int) (*models.Planet, error) {
	return nil, nil
}

func (n *noOpsPersistenceService) SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error) {
	return EmptyPlanetSlice, nil
}

func (n *noOpsPersistenceService) ListAllPlanets(ctx context.Context) ([]*models.Planet, error) {
	return EmptyPlanetSlice, nil
}

func (n *noOpsPersistenceService) RemovePlanetById(ctx context.Context, id int) error {
	return nil
}

func (n *noOpsPersistenceService) RemovePlanetByExactName(ctx context.Context, exactName string) error {
	return nil
}
