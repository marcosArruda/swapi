package services

import (
	"context"

	"github.com/marcosArruda/swapi/pkg/models"
)

type (
	noOpsDatabase struct {
		sm ServiceManager
	}
)

func NewNoOpsDatabase(ctx context.Context) Database {
	return &noOpsDatabase{}
}

func (n *noOpsDatabase) Start(ctx context.Context) error {
	return nil
}

func (n *noOpsDatabase) Close(ctx context.Context) error {
	return nil
}

func (n *noOpsDatabase) Healthy(ctx context.Context) error {
	return nil
}

func (n *noOpsDatabase) WithServiceManager(sm ServiceManager) Database {
	n.sm = sm
	return n
}

func (n *noOpsDatabase) ServiceManager() ServiceManager {
	return n.sm
}

func (n *noOpsDatabase) GetPlanetById(ctx context.Context, id int) (*models.Planet, error) {
	return nil, nil
}

func (n *noOpsDatabase) SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error) {
	return EmptyPlanetSlice, nil
}

func (n *noOpsDatabase) InsertPlanet(ctx context.Context, p *models.Planet) error {
	return nil
}

func (n *noOpsDatabase) UpdatePlanet(ctx context.Context, p *models.Planet) error {
	return nil
}

func (n *noOpsDatabase) ListAllPlanets(ctx context.Context) ([]*models.Planet, error) {
	return EmptyPlanetSlice, nil
}

func (n *noOpsDatabase) RemovePlanetById(ctx context.Context, id int) error {
	return nil
}

func (n *noOpsDatabase) RemovePlanetByExactName(ctx context.Context, exactName string) error {
	return nil
}
