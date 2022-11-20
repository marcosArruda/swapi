package swapiaccess

import (
	"context"
	"reflect"
	"testing"

	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/marcosArruda/swapi/pkg/services"
	"github.com/peterhellberg/swapi"
)

func Test_swApiServiceFinal_Start(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *swApiServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("swApiServiceFinal.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_swApiServiceFinal_Close(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *swApiServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Close(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("swApiServiceFinal.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_swApiServiceFinal_Healthy(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *swApiServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Healthy(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("swApiServiceFinal.Healthy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_swApiServiceFinal_WithServiceManager(t *testing.T) {
	type args struct {
		sm services.ServiceManager
	}
	tests := []struct {
		name string
		n    *swApiServiceFinal
		args args
		want services.SwApiService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.WithServiceManager(tt.args.sm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swApiServiceFinal.WithServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swApiServiceFinal_ServiceManager(t *testing.T) {
	tests := []struct {
		name string
		n    *swApiServiceFinal
		want services.ServiceManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.ServiceManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swApiServiceFinal.ServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swApiServiceFinal_GetPlanetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		n       *swApiServiceFinal
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
				t.Errorf("swApiServiceFinal.GetPlanetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swApiServiceFinal.GetPlanetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swApiServiceFinal_GetFilmById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		n       *swApiServiceFinal
		args    args
		want    *models.Film
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.GetFilmById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("swApiServiceFinal.GetFilmById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swApiServiceFinal.GetFilmById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swApiServiceFinal_SearchPlanetsByName(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		n       *swApiServiceFinal
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
				t.Errorf("swApiServiceFinal.SearchPlanetsByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swApiServiceFinal.SearchPlanetsByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swApiServiceFinal_ToPersistentPlanet(t *testing.T) {
	type args struct {
		ctx    context.Context
		p      *swapi.Planet
		id     int
		expand bool
	}
	tests := []struct {
		name    string
		n       *swApiServiceFinal
		args    args
		want    *models.Planet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.ToPersistentPlanet(tt.args.ctx, tt.args.p, tt.args.id, tt.args.expand)
			if (err != nil) != tt.wantErr {
				t.Errorf("swApiServiceFinal.ToPersistentPlanet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swApiServiceFinal.ToPersistentPlanet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swApiServiceFinal_fillFilmsFromPlanet(t *testing.T) {
	type args struct {
		ctx context.Context
		p   *models.Planet
	}
	tests := []struct {
		name string
		n    *swApiServiceFinal
		args args
		want *models.Planet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.fillFilmsFromPlanet(tt.args.ctx, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swApiServiceFinal.fillFilmsFromPlanet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swApiServiceFinal_ToPersistentFilm(t *testing.T) {
	type args struct {
		ctx    context.Context
		f      *swapi.Film
		id     int
		expand bool
	}
	tests := []struct {
		name    string
		n       *swApiServiceFinal
		args    args
		want    *models.Film
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.ToPersistentFilm(tt.args.ctx, tt.args.f, tt.args.id, tt.args.expand)
			if (err != nil) != tt.wantErr {
				t.Errorf("swApiServiceFinal.ToPersistentFilm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swApiServiceFinal.ToPersistentFilm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swApiServiceFinal_fillPlanetsFromFilm(t *testing.T) {
	type args struct {
		ctx context.Context
		f   *models.Film
	}
	tests := []struct {
		name string
		n    *swApiServiceFinal
		args args
		want *models.Film
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.fillPlanetsFromFilm(tt.args.ctx, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swApiServiceFinal.fillPlanetsFromFilm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swApiServiceFinal_PutOnline(t *testing.T) {
	tests := []struct {
		name string
		n    *swApiServiceFinal
		want services.SwApiService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.PutOnline(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swApiServiceFinal.PutOnline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swApiServiceFinal_PutOffline(t *testing.T) {
	tests := []struct {
		name string
		n    *swApiServiceFinal
		want services.SwApiService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.PutOffline(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swApiServiceFinal.PutOffline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swApiServiceFinal_IsOnline(t *testing.T) {
	tests := []struct {
		name string
		n    *swApiServiceFinal
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.IsOnline(); got != tt.want {
				t.Errorf("swApiServiceFinal.IsOnline() = %v, want %v", got, tt.want)
			}
		})
	}
}
