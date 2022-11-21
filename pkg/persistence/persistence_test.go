package persistence

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

func Test_persistenceServiceFinal_Start(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	sm, ctx := NewManagerForTests()
	ps := NewPersistenceService().WithServiceManager(sm)
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success", // Only success because PersistenceService.Start(ctx) does nothing with the Context for now
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx},
			wantErr: false,
		},
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
	sm, ctx := NewManagerForTests()
	ps := NewPersistenceService().WithServiceManager(sm)
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success", // Only success because PersistenceService.Close(ctx) does nothing with the Context for now
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx},
			wantErr: false,
		},
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
	sm, ctx := NewManagerForTests()
	ps := NewPersistenceService().WithServiceManager(sm)
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success", // Only success because PersistenceService.Healthy(ctx) does nothing with the Context for now
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx},
			wantErr: false,
		},
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
	sm, _ := NewManagerForTests()
	ps := NewPersistenceService()
	tests := []struct {
		name string
		n    *persistenceServiceFinal
		args args
		want services.PersistenceService
	}{
		{
			name: "success",
			n:    ps.(*persistenceServiceFinal),
			args: args{sm: sm},
			want: ps,
		},
		{
			name: "WithNil",
			n:    ps.(*persistenceServiceFinal),
			args: args{sm: nil},
			want: ps,
		},
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
	sm, _ := NewManagerForTests()
	tests := []struct {
		name string
		n    *persistenceServiceFinal
		want services.ServiceManager
	}{
		{
			name: "success",
			n:    NewPersistenceService().WithServiceManager(sm).(*persistenceServiceFinal),
			want: sm,
		},
		{
			name: "returnNil",
			n:    NewPersistenceService().(*persistenceServiceFinal),
			want: nil,
		},
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
	sm, ctx := NewManagerForTests()
	ps := NewPersistenceService().WithServiceManager(sm)
	pp, _ := sm.Database().GetPlanetById(ctx, 1)
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, p: pp},
			wantErr: false,
		},
		{
			name:    "dbGetError",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, p: &models.Planet{Id: 0}},
			wantErr: true,
		},
		{
			name:    "dbGetNil",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, p: &models.Planet{Id: 2}},
			wantErr: true,
		},
		{
			name:    "txBeginError",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: context.WithValue(ctx, "error", struct{}{}), p: &models.Planet{Id: 2}},
			wantErr: true,
		},
		{
			name:    "dbInsertPlanetError",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, p: &models.Planet{Id: 2}},
			wantErr: true,
		},
		{
			name:    "dbInsertPlanetSuccess",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, p: &models.Planet{Id: 3}},
			wantErr: false,
		},
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
	sm, ctx := NewManagerForTests()
	ps := NewPersistenceService().WithServiceManager(sm)
	p := &models.Planet{
		Id:      1,
		Name:    "Terra",
		Climate: "Tropical",
		Terrain: "continental",
		URL:     "https://something.com/api/planet/1/",
	}

	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		want    *models.Planet
		wantErr bool
	}{
		{
			name:    "success",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, id: 1},
			want:    p,
			wantErr: false,
		},
		{
			name:    "error",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, id: 0},
			want:    nil,
			wantErr: true,
		},
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
	sm, ctx := NewManagerForTests()
	ps := NewPersistenceService().WithServiceManager(sm)
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		want    []*models.Planet
		wantErr bool
	}{
		{
			name:    "success",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, name: "somename"},
			want:    services.EmptyPlanetSlice,
			wantErr: false,
		},
		{
			name:    "error",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, name: "error"},
			want:    services.EmptyPlanetSlice,
			wantErr: true,
		},
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
	sm, ctx := NewManagerForTests()
	ps := NewPersistenceService().WithServiceManager(sm)
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		want    []*models.Planet
		wantErr bool
	}{
		{
			name:    "success",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx},
			want:    services.EmptyPlanetSlice,
			wantErr: false,
		},
		{
			name:    "error",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: nil},
			want:    services.EmptyPlanetSlice,
			wantErr: true,
		},
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
	sm, ctx := NewManagerForTests()
	ps := NewPersistenceService().WithServiceManager(sm)
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, id: 1},
			wantErr: false,
		},
		{
			name:    "dbGetError",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, id: 0},
			wantErr: true,
		},
		{
			name:    "dbGetNil",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, id: -1},
			wantErr: false,
		},
		{
			name:    "txBeginError",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: context.WithValue(ctx, "error", struct{}{}), id: 2},
			wantErr: true,
		},
		{
			name:    "dbRemovePlanetError",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: context.WithValue(ctx, "removePlanetError", struct{}{}), id: 1},
			wantErr: true,
		},
		{
			name:    "dbRemovePlanetSuccess",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, id: 1},
			wantErr: false,
		},
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
	sm, ctx := NewManagerForTests()
	ps := NewPersistenceService().WithServiceManager(sm)
	tests := []struct {
		name    string
		n       *persistenceServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: ctx, exactName: "somename"},
			wantErr: false,
		},
		{
			name:    "dbTxBeginError",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: context.WithValue(ctx, "error", struct{}{}), exactName: "empty"},
			wantErr: true,
		},
		{
			name:    "dbRemovePlanetExactNameError",
			n:       ps.(*persistenceServiceFinal),
			args:    args{ctx: context.WithValue(ctx, "removePlanetExactNameError", struct{}{}), exactName: "error"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.RemovePlanetByExactName(tt.args.ctx, tt.args.exactName); (err != nil) != tt.wantErr {
				t.Errorf("persistenceServiceFinal.RemovePlanetByExactName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
