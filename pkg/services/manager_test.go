package services

import (
	"context"
	"reflect"
	"testing"
)

func TestNewManager(t *testing.T) {
	type args struct {
		asyncWorkChannel chan func() error
		stop             chan struct{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success", //just success, since its the "constructor" and there is no ifs inside
			args: args{nil, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewManager(tt.args.asyncWorkChannel, tt.args.stop); !okServiceManager(got) {
				t.Errorf("NewManager() = %v, is Not OK, something is nil", got)
			}
		})
	}
}

func Test_serviceManagerFinal_Start(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	ctx := context.Background()
	sm := NewManager(nil, nil)
	tests := []struct {
		name    string
		m       *serviceManagerFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			m:       sm.(*serviceManagerFinal),
			args:    args{ctx},
			wantErr: false,
		},
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
	ctx := context.Background()
	sm := NewManager(nil, nil)
	tests := []struct {
		name    string
		m       *serviceManagerFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			m:       sm.(*serviceManagerFinal),
			args:    args{ctx},
			wantErr: false,
		},
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
	ctx := context.Background()
	sm := NewManager(nil, nil)
	tests := []struct {
		name    string
		m       *serviceManagerFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			m:       sm.(*serviceManagerFinal),
			args:    args{ctx},
			wantErr: false,
		},
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
	sm := NewManager(nil, nil)
	s := NewNoOpsLogsService()
	tests := []struct {
		name string
		m    *serviceManagerFinal
		args args
		want ServiceManager
	}{
		{
			name: "success",
			m:    sm.(*serviceManagerFinal),
			args: args{s},
			want: sm,
		},
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
	sm := NewManager(nil, nil)
	s := NewNoOpsLogsService()
	sm.WithLogsService(s)
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want LogsService
	}{
		{
			name: "success",
			m:    sm.(*serviceManagerFinal),
			want: s,
		},
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
	sm := NewManager(nil, nil)
	s := NewNoOpsHttpService()
	tests := []struct {
		name string
		m    *serviceManagerFinal
		args args
		want ServiceManager
	}{
		{
			name: "success",
			m:    sm.(*serviceManagerFinal),
			args: args{s},
			want: sm,
		},
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
	sm := NewManager(nil, nil)
	s := NewNoOpsHttpService()
	sm.WithHttpService(s)
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want HttpService
	}{
		{
			name: "success",
			m:    sm.(*serviceManagerFinal),
			want: s,
		},
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
	sm := NewManager(nil, nil)
	s := NewNoOpsPersistenceService()
	tests := []struct {
		name string
		m    *serviceManagerFinal
		args args
		want ServiceManager
	}{
		{
			name: "success",
			m:    sm.(*serviceManagerFinal),
			args: args{s},
			want: sm,
		},
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
	sm := NewManager(nil, nil)
	s := NewNoOpsPersistenceService()
	sm.WithPersistenceService(s)
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want PersistenceService
	}{
		{
			name: "success",
			m:    sm.(*serviceManagerFinal),
			want: s,
		},
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
	sm := NewManager(nil, nil)
	s := NewNoOpsDatabase()
	tests := []struct {
		name string
		m    *serviceManagerFinal
		args args
		want ServiceManager
	}{
		{
			name: "success",
			m:    sm.(*serviceManagerFinal),
			args: args{s},
			want: sm,
		},
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
	sm := NewManager(nil, nil)
	s := NewNoOpsDatabase()
	sm.WithDatabase(s)
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want Database
	}{
		{
			name: "success",
			m:    sm.(*serviceManagerFinal),
			want: s,
		},
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
	sm := NewManager(nil, nil)
	s := NewNoOpsSwService()
	tests := []struct {
		name string
		m    *serviceManagerFinal
		args args
		want ServiceManager
	}{
		{
			name: "success",
			m:    sm.(*serviceManagerFinal),
			args: args{s},
			want: sm,
		},
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
	sm := NewManager(nil, nil)
	s := NewNoOpsSwService()
	sm.WithSwApiService(s)
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want SwApiService
	}{
		{
			name: "success",
			m:    sm.(*serviceManagerFinal),
			want: s,
		},
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
	sm := NewManager(nil, nil)
	s := NewNoOpsPlanetFinderService()
	tests := []struct {
		name string
		m    *serviceManagerFinal
		args args
		want ServiceManager
	}{
		{
			name: "success",
			m:    sm.(*serviceManagerFinal),
			args: args{s},
			want: sm,
		},
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
	sm := NewManager(nil, nil)
	s := NewNoOpsPlanetFinderService()
	sm.WithPlanetFinderService(s)
	tests := []struct {
		name string
		m    *serviceManagerFinal
		want PlanetFinderService
	}{
		{
			name: "success",
			m:    sm.(*serviceManagerFinal),
			want: s,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.PlanetFinderService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceManagerFinal.PlanetFinderService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func okServiceManager(m1 ServiceManager) bool {
	m1T := m1.(*serviceManagerFinal)
	return m1T.database != nil &&
		m1T.httpService != nil &&
		m1T.logsService != nil &&
		m1T.persistenceService != nil &&
		m1T.planetFinderService != nil &&
		m1T.swApiService != nil
}
