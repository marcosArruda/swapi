package swhttp

import (
	"context"

	"github.com/marcosArruda/swapi/pkg/services"
	"github.com/peterhellberg/swapi"
)

type (
	swApiServiceFinal struct {
		sm services.ServiceManager
	}
)

func NewSwService(ctx context.Context) services.SwApiService {
	return &swApiServiceFinal{}
}

func (n *swApiServiceFinal) Start(ctx context.Context) error {
	return nil
}

func (n *swApiServiceFinal) Close(ctx context.Context) error {
	return nil
}

func (n *swApiServiceFinal) Healthy(ctx context.Context) error {
	return nil
}

func (n *swApiServiceFinal) WithServiceManager(sm services.ServiceManager) services.SwApiService {
	n.sm = sm
	return n
}

func (n *swApiServiceFinal) ServiceManager() services.ServiceManager {
	return n.sm
}

func (n *swApiServiceFinal) GetPlanetById(id int) (*swapi.Planet, error) {
	return &swapi.Planet{
		Name: "DummyPlanet",
	}, nil
}
