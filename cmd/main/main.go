package main

import (
	"context"
	"fmt"

	"github.com/marcosArruda/swapi/pkg/httpservice"
	"github.com/marcosArruda/swapi/pkg/logs"
	"github.com/marcosArruda/swapi/pkg/persistence"
	"github.com/marcosArruda/swapi/pkg/planetfinder"
	"github.com/marcosArruda/swapi/pkg/services"
	"github.com/marcosArruda/swapi/pkg/swapiaccess"
)

func main() {
	ctx := context.Background()
	//time.Sleep(5 * time.Second)
	asyncWorkChannel := make(chan func() error)
	stop := make(chan struct{})

	sm := services.NewManager(ctx, asyncWorkChannel, stop).
		WithLogsService(logs.NewLogsService(ctx)).
		WithDatabase(persistence.NewDatabase(ctx)).
		WithPersistenceService(persistence.NewPersistenceService(ctx)).
		WithSwApiService(swapiaccess.NewSwService(ctx).PutOnline()).
		WithPlanetFinderService(planetfinder.NewPlanetFinderService(ctx)).
		WithHttpService(httpservice.NewHttpService(ctx))

	// This is the goroutine that will execute any async work
	go func() {
	basicLoop:
		for {
			select {
			case w := <-asyncWorkChannel:
				if err := w(); err != nil {
					sm.LogsService().Warn(ctx, fmt.Sprintf("an async work failed: %s", err.Error()))
				}
			case <-stop: // triggered when the stop channel is closed
				break basicLoop // stop listening
			}
		}
	}()

	sm.Start(ctx)
	//features
	/*
		- Carregar um planeta da API através do Id
		- Listar os planetas
		- Buscar planeta por nome
		- Buscar por ID
		- Remover planeta

		*** PARA CADA PLANETA ***
			- Nome
			- clima
			- terreno
			- filmes
		*** PARA CADA FILME ***
			- nome
			- diretor
			- data de lançamento
	*/
}
