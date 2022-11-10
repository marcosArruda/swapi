package main

import (
	"context"

	"github.com/marcosArruda/swapi/pkg/logs"
	"github.com/marcosArruda/swapi/pkg/services"
)

func main() {
	ctx := context.Background()
	sm := services.NewManager(ctx).
		WithLogsService(logs.NewLogsService(ctx)).
		WithHttpService(services.NewHttpService(ctx))
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
