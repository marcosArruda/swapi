package swapiaccess

import (
	"context"

	"github.com/marcosArruda/swapi/pkg/messages"
	"github.com/marcosArruda/swapi/pkg/services"
	"github.com/peterhellberg/swapi"
)

type (
	swApiServiceFinal struct {
		sm       services.ServiceManager
		swclient *swapi.Client
		online   bool
	}
)

func NewSwService(ctx context.Context, online bool) services.SwApiService {
	return &swApiServiceFinal{swclient: swapi.DefaultClient, online: online}
}

func (n *swApiServiceFinal) Start(ctx context.Context) error {
	n.sm.LogsService().Info(ctx, "SwApi Service Started Started!")
	return nil
}

func (n *swApiServiceFinal) Close(ctx context.Context) error {
	return nil
}

func (n *swApiServiceFinal) Healthy(ctx context.Context) error {
	if n.online {
		_, err := n.swclient.Planet(1)
		if err != nil {
			n.sm.LogsService().Warn(ctx, messages.SwApiUnavailableError.Error())
			n.online = false
			return messages.SwApiUnavailableError
		}
	}
	return nil
}

func (n *swApiServiceFinal) WithServiceManager(sm services.ServiceManager) services.SwApiService {
	n.sm = sm
	return n
}

func (n *swApiServiceFinal) ServiceManager() services.ServiceManager {
	return n.sm
}

func (n *swApiServiceFinal) GetPlanetById(ctx context.Context, id int) (*swapi.Planet, error) {
	if n.online {
		p, err := n.swclient.Planet(id)
		if err != nil {
			return nil, messages.SwApiUnavailableError
		}
		return &p, nil
	}
	return nil, messages.SwApiIsOfflineError
}

func (n *swApiServiceFinal) PutOnline() {
	n.online = true
}

func (n *swApiServiceFinal) PutOffline() {
	n.online = false
}

func (n *swApiServiceFinal) IsOnline() bool {
	return n.online
}
