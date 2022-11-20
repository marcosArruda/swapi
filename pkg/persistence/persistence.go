package persistence

import (
	"context"
	"fmt"

	"github.com/marcosArruda/swapi/pkg/messages"
	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/marcosArruda/swapi/pkg/services"
)

type (
	persistenceServiceFinal struct {
		sm services.ServiceManager
	}
)

func NewPersistenceService(ctx context.Context) services.PersistenceService {
	return &persistenceServiceFinal{}
}

func (n *persistenceServiceFinal) Start(ctx context.Context) error {
	n.sm.LogsService().Info(ctx, "Persistence Started!")
	return nil
}

func (n *persistenceServiceFinal) Close(ctx context.Context) error {
	return nil
}

func (n *persistenceServiceFinal) Healthy(ctx context.Context) error {
	return nil
}

func (n *persistenceServiceFinal) WithServiceManager(sm services.ServiceManager) services.PersistenceService {
	n.sm = sm
	return n
}

func (n *persistenceServiceFinal) ServiceManager() services.ServiceManager {
	return n.sm
}

func (n *persistenceServiceFinal) UpsertPlanet(ctx context.Context, p *models.Planet) error {
	db := n.ServiceManager().Database()
	n.sm.LogsService().Info(ctx, fmt.Sprintf("Upserting planet {id: %d, name: %s}", p.Id, p.Name))
	pp, err := db.GetPlanetById(ctx, p.Id)
	if err != nil && err != messages.NoPlanetFound {
		errCustom := &messages.PlanetError{
			Msg: fmt.Sprintf("Error Upserting planet named '%s' with ID '%d': %s", p.Name, p.Id, err.Error())}
		n.ServiceManager().LogsService().Error(ctx, errCustom.Msg)
		return errCustom
	}
	if pp != nil {
		return db.UpdatePlanet(ctx, p)
	}
	tx, err := db.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	err = db.InsertPlanet(ctx, tx, p)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (n *persistenceServiceFinal) GetPlanetById(ctx context.Context, id int) (*models.Planet, error) {
	return n.ServiceManager().Database().GetPlanetById(ctx, id)
}

func (n *persistenceServiceFinal) SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error) {
	return n.sm.Database().SearchPlanetsByName(ctx, name)
}

func (n *persistenceServiceFinal) ListAllPlanets(ctx context.Context) ([]*models.Planet, error) {
	return n.sm.Database().ListAllPlanets(ctx)
}

func (n *persistenceServiceFinal) RemovePlanetById(ctx context.Context, id int) error {
	db := n.ServiceManager().Database()
	p, err := db.GetPlanetById(ctx, id)
	if err != nil && err != messages.NoPlanetFound {
		errCustom := &messages.PlanetError{
			Msg: fmt.Sprintf("Error Removing planet named with ID '%d': %s", p.Id, err.Error())}
		n.ServiceManager().LogsService().Error(ctx, errCustom.Msg)
		return errCustom
	} else if err == messages.NoPlanetFound {
		n.ServiceManager().LogsService().Info(ctx, fmt.Sprintf("Planet with ID '%d' are not yet in the local database. Please Load it with GET /planet/%d first!", id, id))
		return nil
	}
	tx, err := db.BeginTransaction(ctx)

	if err != nil {
		n.ServiceManager().LogsService().Error(ctx, fmt.Sprintf("Error opening transaction for removing the planet with ID '%d': %s", id, err.Error()))
		db.RollbackTransaction(tx)
		return err
	}

	if err = n.sm.Database().RemovePlanetById(ctx, tx, id); err != nil {
		n.ServiceManager().LogsService().Error(ctx, fmt.Sprintf("Error Deleting Planet with ID '%d': %s", id, err.Error()))
		db.RollbackTransaction(tx)
		return err
	}
	return db.CommitTransaction(tx)
}

func (n *persistenceServiceFinal) RemovePlanetByExactName(ctx context.Context, exactName string) error {
	db := n.ServiceManager().Database()
	tx, err := db.BeginTransaction(ctx)
	if err != nil {
		n.ServiceManager().LogsService().Error(ctx, fmt.Sprintf("Error opening transaction for removing the planet with name '%s': %s", exactName, err.Error()))
		db.RollbackTransaction(tx)
		return err
	}
	if err = n.ServiceManager().Database().RemovePlanetByExactName(ctx, tx, exactName); err != nil {
		n.ServiceManager().LogsService().Error(ctx, fmt.Sprintf("Error Removing planet by exact name '%s': %s", exactName, err.Error()))
		db.RollbackTransaction(tx)
		return err
	}
	return db.CommitTransaction(tx)
}
