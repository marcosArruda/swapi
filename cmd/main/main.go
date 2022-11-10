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
}
