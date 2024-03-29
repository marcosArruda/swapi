package planetfinder

import (
	"context"
	"fmt"
	"time"

	"github.com/marcosArruda/swapi/pkg/messages"
	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/marcosArruda/swapi/pkg/services"
)

type (
	PlanetNameTag            string
	planetFinderServiceFinal struct {
		sm services.ServiceManager
	}
)

var (
	PlanetNameTagLabel PlanetNameTag = "PlanetNameTag"
)

func NewPlanetFinderService() services.PlanetFinderService {
	return &planetFinderServiceFinal{}
}

func (n *planetFinderServiceFinal) Start(ctx context.Context) error {
	n.sm.LogsService().Info(ctx, "PlanetFinder Service Started!")
	return nil
}

func (n *planetFinderServiceFinal) Close(ctx context.Context) error {
	return nil
}

func (n *planetFinderServiceFinal) Healthy(ctx context.Context) error {
	return nil
}

func (n *planetFinderServiceFinal) WithServiceManager(sm services.ServiceManager) services.PlanetFinderService {
	n.sm = sm
	return n
}

func (n *planetFinderServiceFinal) ServiceManager() services.ServiceManager {
	return n.sm
}

func (n *planetFinderServiceFinal) GetPlanetById(ctx context.Context, id int) (*models.Planet, error) {
	p, err := n.sm.PersistenceService().GetPlanetById(ctx, id)
	if err != nil && err != messages.NoPlanetFound {
		return nil, err
	}
	if p != nil {
		return p, nil
	}

	if n.sm.SwApiService().IsOnline() {
		pp, err := n.sm.SwApiService().GetPlanetById(ctx, id)
		if err != nil {
			msg := fmt.Sprintf("System is in 'online' mode but we can't find any Planet with ID %d on the Public SW API: %s", id, err.Error())
			n.sm.LogsService().Error(ctx, msg)
			return nil, &messages.PlanetError{
				PlanetId: id,
				Msg:      msg,
			}
		}
		asynContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		go func() {
			n.sm.AsyncWorkChannel() <- func() error { //async persistence ...
				defer cancel()
				if err = n.sm.PersistenceService().UpsertPlanet(asynContext, pp); err != nil {
					return err
				}
				return nil
			}
		}()

		return pp, nil
	}
	return nil, &messages.PlanetError{
		PlanetId: id,
		Msg:      fmt.Sprintf("System is in 'offline' mode and there is no Planet with ID %d", id),
	}
}

func (n *planetFinderServiceFinal) SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error) {
	ps, err := n.sm.SwApiService().SearchPlanetsByName(ctx, name)
	if err != nil {
		return services.EmptyPlanetSlice, err
	}

	go func() {
		for _, pp := range ps {
			asynContext := context.WithValue(context.Background(), PlanetNameTagLabel, pp.Name)
			n.sm.AsyncWorkChannel() <- func() error { //async persistence ...
				if err = n.sm.PersistenceService().UpsertPlanet(asynContext, pp); err != nil {
					return err
				}
				return nil
			}
		}
	}()
	return ps, nil
}

func (n *planetFinderServiceFinal) ListAllPlanets(ctx context.Context) ([]*models.Planet, error) {
	return n.sm.PersistenceService().ListAllPlanets(ctx)
}

func (n *planetFinderServiceFinal) RemovePlanetById(ctx context.Context, id int) error {
	return n.sm.PersistenceService().RemovePlanetById(ctx, id)
}

func (n *planetFinderServiceFinal) RemovePlanetByExactName(ctx context.Context, exactName string) error {
	return n.sm.PersistenceService().RemovePlanetByExactName(ctx, exactName)
}
