package services

import (
	"context"
	"reflect"
	"testing"
)

func TestNewManager(t *testing.T) {
	type args struct {
		ctx              context.Context
		asyncWorkChannel chan func() error
		stop             chan struct{}
	}
	tests := []struct {
		name string
		args args
		want ServiceManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewManager(tt.args.ctx, tt.args.asyncWorkChannel, tt.args.stop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_Start(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		m       *serviceManagerFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("serviceManagerFinal.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_serviceManagerFinal_Close(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		m       *serviceManagerFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Close(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("serviceManagerFinal.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_serviceManagerFinal_Healthy(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		m       *serviceManagerFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Healthy(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("serviceManagerFinal.Healthy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_serviceManagerFinal_WithLogsService(t *testing.T) {
	type args struct {
		ls LogsService
	}
	tests := []struct {
		name string
		m    *serviceManagerFinal
		args args
		want ServiceManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.WithLogsService(tt.args.ls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.WithLogsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_LogsService(t *testing.T) {
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want LogsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.LogsService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.LogsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_WithHttpService(t *testing.T) {
	type args struct {
		h HttpService
	}
	tests := []struct {
		name string
		m    *serviceManagerFinal
		args args
		want ServiceManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.WithHttpService(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.WithHttpService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_HttpService(t *testing.T) {
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want HttpService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.HttpService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.HttpService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_WithPersistenceService(t *testing.T) {
	type args struct {
		p PersistenceService
	}
	tests := []struct {
		name string
		m    *serviceManagerFinal
		args args
		want ServiceManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.WithPersistenceService(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.WithPersistenceService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_PersistenceService(t *testing.T) {
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want PersistenceService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.PersistenceService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.PersistenceService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_WithDatabase(t *testing.T) {
	type args struct {
		db Database
	}
	tests := []struct {
		name string
		m    *serviceManagerFinal
		args args
		want ServiceManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.WithDatabase(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.WithDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_Database(t *testing.T) {
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want Database
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Database(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.Database() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_WithSwApiService(t *testing.T) {
	type args struct {
		sw SwApiService
	}
	tests := []struct {
		name string
		m    *serviceManagerFinal
		args args
		want ServiceManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.WithSwApiService(tt.args.sw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.WithSwApiService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_SwApiService(t *testing.T) {
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want SwApiService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.SwApiService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.SwApiService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_WithPlanetFinderService(t *testing.T) {
	type args struct {
		p PlanetFinderService
	}
	tests := []struct {
		name string
		m    *serviceManagerFinal
		args args
		want ServiceManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.WithPlanetFinderService(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.WithPlanetFinderService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_PlanetFinderService(t *testing.T) {
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want PlanetFinderService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.PlanetFinderService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.PlanetFinderService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceManagerFinal_AsyncWorkChannel(t *testing.T) {
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want chan func() error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.AsyncWorkChannel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.AsyncWorkChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}
