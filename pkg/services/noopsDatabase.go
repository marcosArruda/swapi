package services

import (
	"context"
	"database/sql"

	"github.com/marcosArruda/swapi/pkg/models"
)

type (
	noOpsDatabase struct {
		sm ServiceManager
	}
)

func NewNoOpsDatabase() Database {
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

func (n *noOpsDatabase) BeginTransaction(ctx context.Context) (*sql.Tx, error) {
	return nil, nil
}
func (n *noOpsDatabase) CommitTransaction(tx *sql.Tx) error {
	return nil
}

func (n *noOpsDatabase) RollbackTransaction(tx *sql.Tx) error {
	return nil
}

func (n *noOpsDatabase) GetPlanetById(ctx context.Context, id int) (*models.Planet, error) {
	return nil, nil
}

func (n *noOpsDatabase) SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error) {
	return EmptyPlanetSlice, nil
}

func (n *noOpsDatabase) InsertPlanet(ctx context.Context, tx *sql.Tx, p *models.Planet) error {
	return nil
}

func (n *noOpsDatabase) UpdatePlanet(ctx context.Context, p *models.Planet) error {
	return nil
}

func (n *noOpsDatabase) ListAllPlanets(ctx context.Context) ([]*models.Planet, error) {
	return EmptyPlanetSlice, nil
}

func (n *noOpsDatabase) RemovePlanetById(ctx context.Context, tx *sql.Tx, id int) error {
	return nil
}

func (n *noOpsDatabase) RemovePlanetByExactName(ctx context.Context, tx *sql.Tx, exactName string) error {
	return nil
}
