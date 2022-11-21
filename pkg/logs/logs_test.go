package logs

import (
	"context"
	"reflect"
	"testing"

	"github.com/marcosArruda/swapi/pkg/services"
)

func TestNewLogsService(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want services.LogsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogsService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logsServiceFinal_Start(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		f       *logsServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name    string
		f       *logsServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name    string
		f       *logsServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name string
		f    *logsServiceFinal
		args args
		want services.LogsService
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name string
		f    *logsServiceFinal
		want services.ServiceManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ServiceManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("logsServiceFinal.ServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logsServiceFinal_Info(t *testing.T) {
	type args struct {
		ctx context.Context
		s   string
	}
	tests := []struct {
		name string
		f    *logsServiceFinal
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.Info(tt.args.ctx, tt.args.s)
		})
	}
}

func Test_logsServiceFinal_Warn(t *testing.T) {
	type args struct {
		ctx context.Context
		s   string
	}
	tests := []struct {
		name string
		f    *logsServiceFinal
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.Warn(tt.args.ctx, tt.args.s)
		})
	}
}

func Test_logsServiceFinal_Error(t *testing.T) {
	type args struct {
		ctx context.Context
		s   string
	}
	tests := []struct {
		name string
		f    *logsServiceFinal
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.Error(tt.args.ctx, tt.args.s)
		})
	}
}

func Test_logsServiceFinal_Debug(t *testing.T) {
	type args struct {
		ctx context.Context
		s   string
	}
	tests := []struct {
		name string
		f    *logsServiceFinal
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.Debug(tt.args.ctx, tt.args.s)
		})
	}
}
