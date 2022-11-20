package persistence

import (
	"context"
	"reflect"
	"testing"

	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/marcosArruda/swapi/pkg/services"
)

func Test_persistenceServiceFinal_Start(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("persistenceServiceFinal.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_persistenceServiceFinal_Close(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Close(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("persistenceServiceFinal.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_persistenceServiceFinal_Healthy(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Healthy(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("persistenceServiceFinal.Healthy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_persistenceServiceFinal_WithServiceManager(t *testing.T) {
	type args struct {
		sm services.ServiceManager
	}
	tests := []struct {
		name string
		n    *persistenceServiceFinal
		args args
		want services.PersistenceService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.WithServiceManager(tt.args.sm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("persistenceServiceFinal.WithServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_persistenceServiceFinal_ServiceManager(t *testing.T) {
	tests := []struct {
		name string
		n    *persistenceServiceFinal
		want services.ServiceManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.ServiceManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("persistenceServiceFinal.ServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_persistenceServiceFinal_UpsertPlanet(t *testing.T) {
	type args struct {
		ctx context.Context
		p   *models.Planet
	}
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.UpsertPlanet(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("persistenceServiceFinal.UpsertPlanet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_persistenceServiceFinal_GetPlanetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		want    *models.Planet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.GetPlanetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("persistenceServiceFinal.GetPlanetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("persistenceServiceFinal.GetPlanetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_persistenceServiceFinal_SearchPlanetsByName(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		want    []*models.Planet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.SearchPlanetsByName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("persistenceServiceFinal.SearchPlanetsByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("persistenceServiceFinal.SearchPlanetsByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_persistenceServiceFinal_ListAllPlanets(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		want    []*models.Planet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.ListAllPlanets(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("persistenceServiceFinal.ListAllPlanets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("persistenceServiceFinal.ListAllPlanets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_persistenceServiceFinal_RemovePlanetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.RemovePlanetById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("persistenceServiceFinal.RemovePlanetById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_persistenceServiceFinal_RemovePlanetByExactName(t *testing.T) {
	type args struct {
		ctx       context.Context
		exactName string
	}
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.RemovePlanetByExactName(tt.args.ctx, tt.args.exactName); (err != nil) != tt.wantErr {
				t.Errorf("persistenceServiceFinal.RemovePlanetByExactName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
