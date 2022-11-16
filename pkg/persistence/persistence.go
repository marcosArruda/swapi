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
	pp, err := db.GetPlanetById(ctx, p.Id)
	if err != nil && err != messages.NoPlanetFound {
		errCustom := &messages.PlanetError{
			Msg: fmt.Sprintf("Error Upserting planet named '%s' with ID '%d': %s", p.Name, p.Id, err.Error())}
		n.ServiceManager().LogsService().Error(ctx, errCustom.Msg)
		return err
	}
	if pp != nil {
		return db.UpdatePlanet(ctx, p)
	}
	return db.InsertPlanet(ctx, p)
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
	return n.sm.Database().RemovePlanetById(ctx, id)
}

func (n *persistenceServiceFinal) RemovePlanetByExactName(ctx context.Context, exactName string) error {
	return nil
}
