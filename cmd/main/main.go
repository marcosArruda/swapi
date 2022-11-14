package main

import (
	"context"

	"github.com/marcosArruda/swapi/pkg/httpservice"
	"github.com/marcosArruda/swapi/pkg/logs"
	"github.com/marcosArruda/swapi/pkg/services"
)

func main() {
	ctx := context.Background()
	services.NewManager(ctx).
		WithLogsService(logs.NewLogsService(ctx)).
		//WithDatabase(persistence.NewDatabase(ctx)).
		//WithPersistenceService(persistence.NewPersistenceService(ctx)).
		//WithSwApiService(swapiaccess.NewSwService(ctx, true)).
		//WithPlanetFinderService(planetfinder.NewPlanetFinderService(ctx)).
		WithHttpService(httpservice.NewHttpService(ctx)).
		Start(ctx)

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
