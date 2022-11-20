package planetfinder

import (
	"context"
	"reflect"
	"testing"

	"github.com/marcosArruda/swapi/pkg/logs"
	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/marcosArruda/swapi/pkg/services"
)

func NewManagerForTests() (services.ServiceManager, context.Context) {
	asyncWorkChannel := make(chan func() error)
	stop := make(chan struct{})
	ctx := context.Background()
	ctx = context.WithValue(ctx, logs.AppEnvKey, "TESTS")
	ctx = context.WithValue(ctx, logs.AppNameKey, logs.AppName)
	ctx = context.WithValue(ctx, logs.AppVersionKey, logs.AppVersion)
	return services.NewManager(asyncWorkChannel, stop), ctx
}

func Test_planetFinderServiceFinal_Start(t *testing.T) {

	type args struct {
		ctx context.Context
	}
	_, ctx := NewManagerForTests()
	tests := []struct {
		name    string
		n       *planetFinderServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success", //here is just the success case is needed because PlanetFinderService.Start(ctx) does nothing with the ctx yet
			args:    args{ctx: ctx},
			n:       NewPlanetFinderService().(*planetFinderServiceFinal),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("planetFinderServiceFinal.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_planetFinderServiceFinal_Close(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	_, ctx := NewManagerForTests()
	tests := []struct {
		name    string
		n       *planetFinderServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success", //here is just the success case is needed because PlanetFinderService.Close(ctx) does nothing with the ctx yet
			args:    args{ctx: ctx},
			n:       NewPlanetFinderService().(*planetFinderServiceFinal),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Close(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("planetFinderServiceFinal.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_planetFinderServiceFinal_Healthy(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	_, ctx := NewManagerForTests()
	tests := []struct {
		name    string
		n       *planetFinderServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success", //here is just the success case is needed because PlanetFinderService.Healthy(ctx) does nothing with the ctx yet
			args:    args{ctx: ctx},
			n:       NewPlanetFinderService().(*planetFinderServiceFinal),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Healthy(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("planetFinderServiceFinal.Healthy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_planetFinderServiceFinal_WithServiceManager(t *testing.T) {
	type args struct {
		sm services.ServiceManager
	}
	sm, _ := NewManagerForTests()
	pf := NewPlanetFinderService()
	tests := []struct {
		name string
		n    *planetFinderServiceFinal
		args args
		want services.PlanetFinderService
	}{
		{
			name: "success",
			args: args{sm: sm},
			n:    pf.(*planetFinderServiceFinal),
			want: pf,
		},
		{
			name: "successNil",
			args: args{sm: nil},
			n:    pf.(*planetFinderServiceFinal),
			want: pf,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.WithServiceManager(tt.args.sm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("planetFinderServiceFinal.WithServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_planetFinderServiceFinal_ServiceManager(t *testing.T) {
	sm, _ := NewManagerForTests()
	pf := NewPlanetFinderService()
	tests := []struct {
		name string
		n    *planetFinderServiceFinal
		want services.ServiceManager
	}{
		{
			name: "success",
			n:    pf.WithServiceManager(sm).(*planetFinderServiceFinal),
			want: sm,
		},
		{
			name: "successNil",
			n:    pf.WithServiceManager(nil).(*planetFinderServiceFinal),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.ServiceManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("planetFinderServiceFinal.ServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_planetFinderServiceFinal_GetPlanetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		n       *planetFinderServiceFinal
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
				t.Errorf("planetFinderServiceFinal.GetPlanetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("planetFinderServiceFinal.GetPlanetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_planetFinderServiceFinal_SearchPlanetsByName(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		n       *planetFinderServiceFinal
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
				t.Errorf("planetFinderServiceFinal.SearchPlanetsByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("planetFinderServiceFinal.SearchPlanetsByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_planetFinderServiceFinal_ListAllPlanets(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *planetFinderServiceFinal
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
				t.Errorf("planetFinderServiceFinal.ListAllPlanets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("planetFinderServiceFinal.ListAllPlanets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_planetFinderServiceFinal_RemovePlanetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		n       *planetFinderServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.RemovePlanetById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("planetFinderServiceFinal.RemovePlanetById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_planetFinderServiceFinal_RemovePlanetByExactName(t *testing.T) {
	type args struct {
		ctx       context.Context
		exactName string
	}
	tests := []struct {
		name    string
		n       *planetFinderServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.RemovePlanetByExactName(tt.args.ctx, tt.args.exactName); (err != nil) != tt.wantErr {
				t.Errorf("planetFinderServiceFinal.RemovePlanetByExactName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
