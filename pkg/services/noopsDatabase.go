package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marcosArruda/swapi/pkg/messages"
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
	if ctx != nil {
		if ctx.Value("error") != nil {
			return nil, errors.New("some error")
		}
		return &sql.Tx{}, nil
	}
	return nil, nil
}
func (n *noOpsDatabase) CommitTransaction(tx *sql.Tx) error {
	if tx != nil {
		return nil
	}
	return errors.New("some error")
}

func (n *noOpsDatabase) RollbackTransaction(tx *sql.Tx) error {
	if tx != nil {
		return nil
	}
	return errors.New("some error")
}

func (n *noOpsDatabase) GetPlanetById(ctx context.Context, id int) (*models.Planet, error) {
	if id == -1 {
		return nil, messages.NoPlanetFound
	}
	if id == 1 {
		return &models.Planet{
			Id:      1,
			Name:    "Terra",
			Climate: "Tropical",
			Terrain: "continental",
			URL:     "https://something.com/api/planet/1/",
		}, nil
	} else if id == 2 || id == 3 {
		return nil, nil
	}
	return nil, errors.New("some error")
}

func (n *noOpsDatabase) SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error) {
	if name == "error" {
		return EmptyPlanetSlice, errors.New("some error")
	}
	return EmptyPlanetSlice, nil
}

func (n *noOpsDatabase) InsertPlanet(ctx context.Context, tx *sql.Tx, p *models.Planet) error {
	if p.Id == 2 {
		return errors.New("some error")
	}
	return nil
}

func (n *noOpsDatabase) UpdatePlanet(ctx context.Context, p *models.Planet) error {
	return nil
}

func (n *noOpsDatabase) ListAllPlanets(ctx context.Context) ([]*models.Planet, error) {
	if ctx == nil {
		return EmptyPlanetSlice, errors.New("some error")
	}
	return EmptyPlanetSlice, nil
}

func (n *noOpsDatabase) RemovePlanetById(ctx context.Context, tx *sql.Tx, id int) error {
	if ctx != nil && ctx.Value("removePlanetError") != nil {
		return errors.New("some error")
	}
	return nil
}

func (n *noOpsDatabase) RemovePlanetByExactName(ctx context.Context, tx *sql.Tx, exactName string) error {
	if ctx != nil && ctx.Value("removePlanetExactNameError") != nil {
		return errors.New("some error")
	}
	return nil
}
