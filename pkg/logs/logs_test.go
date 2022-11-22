package logs

import (
	"context"
	"reflect"
	"testing"

	"github.com/marcosArruda/swapi/pkg/services"
)

func NewManagerForTests() (services.ServiceManager, context.Context) {
	asyncWorkChannel := make(chan func() error)
	stop := make(chan struct{})
	ctx := context.Background()
	ctx = context.WithValue(ctx, AppEnvKey, "TESTS")
	ctx = context.WithValue(ctx, AppNameKey, AppName)
	ctx = context.WithValue(ctx, AppVersionKey, AppVersion)
	return services.NewManager(asyncWorkChannel, stop), ctx
}

func TestNewLogsService(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogsService(); got == nil {
				t.Errorf("Got a Nil Logger!")
			}
		})
	}
}

func Test_logsServiceFinal_Start(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	_, ctx := NewManagerForTests()
	s := NewLogsService()
	tests := []struct {
		name    string
		f       *logsServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			f:       s.(*logsServiceFinal),
			args:    args{ctx},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("logsServiceFinal.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_logsServiceFinal_Close(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	_, ctx := NewManagerForTests()
	s := NewLogsService()
	tests := []struct {
		name    string
		f       *logsServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			f:       s.(*logsServiceFinal),
			args:    args{ctx},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.Close(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("logsServiceFinal.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_logsServiceFinal_Healthy(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	_, ctx := NewManagerForTests()
	s := NewLogsService()
	tests := []struct {
		name    string
		f       *logsServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			f:       s.(*logsServiceFinal),
			args:    args{ctx},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.Healthy(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("logsServiceFinal.Healthy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_logsServiceFinal_WithServiceManager(t *testing.T) {
	type args struct {
		sm services.ServiceManager
	}
	sm, _ := NewManagerForTests()
	s := NewLogsService()
	tests := []struct {
		name string
		f    *logsServiceFinal
		args args
		want services.LogsService
	}{
		{
			name: "success",
			f:    s.(*logsServiceFinal),
			args: args{sm},
			want: s,
		},
		{
			name: "successnil",
			f:    s.(*logsServiceFinal),
			args: args{sm: nil},
			want: s,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.WithServiceManager(tt.args.sm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("logsServiceFinal.WithServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logsServiceFinal_ServiceManager(t *testing.T) {
	sm, _ := NewManagerForTests()
	s := NewLogsService()
	tests := []struct {
		name string
		f    *logsServiceFinal
		want services.ServiceManager
	}{
		{
			name: "success",
			f:    sm.WithLogsService(s).LogsService().(*logsServiceFinal),
			want: sm,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ServiceManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("logsServiceFinal.ServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
