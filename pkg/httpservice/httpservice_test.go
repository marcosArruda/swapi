package httpservice

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcosArruda/swapi/pkg/logs"
	"github.com/marcosArruda/swapi/pkg/services"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func NewManagerForTests() (services.ServiceManager, context.Context) {
	asyncWorkChannel := make(chan func() error)
	stop := make(chan struct{})
	ctx := context.Background()
	ctx = context.WithValue(ctx, logs.AppEnvKey, "TESTS")
	ctx = context.WithValue(ctx, logs.AppNameKey, logs.AppName)
	ctx = context.WithValue(ctx, logs.AppVersionKey, logs.AppVersion)
	return services.NewManager(asyncWorkChannel, stop), ctx
}

func NewGinContextForTests(reqPath string) *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	u := &url.URL{
		Path: reqPath,
	}
	v := url.Values{}
	v.Add("search", "somesearch")
	v.Add("name", "somename")
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    u,
	}

	ctx.Request.URL.RawQuery = v.Encode()
	return ctx
}
func Test_httpServiceFinal_Start(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	sm, ctx := NewManagerForTests()
	httpService := sm.WithHttpService(NewHttpService()).HttpService()
	tests := []struct {
		name    string
		n       *httpServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			n:       httpService.(*httpServiceFinal),
			args:    args{ctx},
			wantErr: false,
		},
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
	sm, ctx := NewManagerForTests()
	httpService := sm.WithHttpService(NewHttpService()).HttpService()
	httpService.Start(ctx)
	tests := []struct {
		name    string
		n       *httpServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			n:       httpService.(*httpServiceFinal),
			args:    args{ctx},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Close(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("httpServiceFinal.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_httpServiceFinal_GetPlanetById(t *testing.T) {
	type args struct {
		c *gin.Context
	}

	sm, _ := NewManagerForTests()
	httpService := sm.WithHttpService(NewHttpService()).HttpService()
	ginCtx := NewGinContextForTests("/some-request-path/1/")
	tests := []struct {
		name string
		n    *httpServiceFinal
		args args
	}{
		{
			name: "success",
			n:    httpService.(*httpServiceFinal),
			args: args{ginCtx},
		},
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
	sm, _ := NewManagerForTests()
	httpService := sm.WithHttpService(NewHttpService()).HttpService()
	ginCtx := NewGinContextForTests("/some-request-path?search=somename")
	tests := []struct {
		name string
		n    *httpServiceFinal
		args args
	}{
		{
			name: "success",
			n:    httpService.(*httpServiceFinal),
			args: args{ginCtx},
		},
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
	sm, _ := NewManagerForTests()
	httpService := sm.WithHttpService(NewHttpService()).HttpService()
	ginCtx := NewGinContextForTests("/some-request-path/1/")
	tests := []struct {
		name string
		n    *httpServiceFinal
		args args
	}{
		{
			name: "success",
			n:    httpService.(*httpServiceFinal),
			args: args{ginCtx},
		},
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
	sm, _ := NewManagerForTests()
	httpService := sm.WithHttpService(NewHttpService()).HttpService()
	ginCtx := NewGinContextForTests("/some-request-path/1/")
	tests := []struct {
		name string
		n    *httpServiceFinal
		args args
	}{
		{
			name: "success",
			n:    httpService.(*httpServiceFinal),
			args: args{ginCtx},
		},
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
	sm, _ := NewManagerForTests()
	httpService := sm.WithHttpService(NewHttpService()).HttpService()
	ginCtx := NewGinContextForTests("/some-request-path/?name=somename")
	tests := []struct {
		name string
		n    *httpServiceFinal
		args args
	}{
		{
			name: "success",
			n:    httpService.(*httpServiceFinal),
			args: args{ginCtx},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.RemovePlanetByExactName(tt.args.c)
		})
	}
}
