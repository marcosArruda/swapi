package httpservice

import (
	"context"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcosArruda/swapi/pkg/services"
)

func Test_httpServiceFinal_Start(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *httpServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("httpServiceFinal.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_httpServiceFinal_Close(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *httpServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Close(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("httpServiceFinal.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_httpServiceFinal_Healthy(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		n       *httpServiceFinal
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Healthy(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("httpServiceFinal.Healthy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_httpServiceFinal_WithServiceManager(t *testing.T) {
	type args struct {
		sm services.ServiceManager
	}
	tests := []struct {
		name string
		n    *httpServiceFinal
		args args
		want services.HttpService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.WithServiceManager(tt.args.sm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpServiceFinal.WithServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpServiceFinal_ServiceManager(t *testing.T) {
	tests := []struct {
		name string
		n    *httpServiceFinal
		want services.ServiceManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.ServiceManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpServiceFinal.ServiceManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpServiceFinal_GetPlanetById(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		n    *httpServiceFinal
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.GetPlanetById(tt.args.c)
		})
	}
}

func Test_httpServiceFinal_SearchPlanetsByName(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		n    *httpServiceFinal
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.SearchPlanetsByName(tt.args.c)
		})
	}
}

func Test_httpServiceFinal_ListAllPlanets(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		n    *httpServiceFinal
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.ListAllPlanets(tt.args.c)
		})
	}
}

func Test_httpServiceFinal_RemovePlanetById(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		n    *httpServiceFinal
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.RemovePlanetById(tt.args.c)
		})
	}
}

func Test_httpServiceFinal_RemovePlanetByExactName(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		n    *httpServiceFinal
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.RemovePlanetByExactName(tt.args.c)
		})
	}
}
