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
	pp, err := n.ServiceManager().Database().GetPlanetById(ctx, p.Id)
	if err != nil {
		errCustom := &messages.PlanetError{
			Msg: fmt.Sprintf("Error Upserting planet named '%s' with ID '%d': %s", p.Name, p.Id, err.Error())}
		n.ServiceManager().LogsService().Error(errCustom.Msg)
		return err
	}
	if pp == nil {

	}
	return nil
}
