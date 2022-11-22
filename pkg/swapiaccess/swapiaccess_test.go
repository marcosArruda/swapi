package swapiaccess

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/marcosArruda/swapi/pkg/logs"
	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/marcosArruda/swapi/pkg/services"
	"github.com/peterhellberg/swapi"
)

type (
	searchableHttpClientMock struct {
	}
	swApiClientMock struct {
		dropError bool
	}
)

var (
	basicPlanet = &models.Planet{
		Id:      1,
		Name:    "Terra",
		Climate: "tropical",
		Terrain: "terra",
		Films: []*models.Film{
			{
				Id:        1,
				Title:     "Filme da terra",
				EpisodeID: 1,
				Created:   "800 quintilhões de anos atras",
				Director:  "Único",
				URL:       "https://something.com/api/film/1/",
			},
		},
		URL: "https://something.com/api/planet/1/",
	}
)

func (s *swApiClientMock) Vehicle(id int) (swapi.Vehicle, error)   { return swapi.Vehicle{}, nil }
func (s *swApiClientMock) Starship(id int) (swapi.Starship, error) { return swapi.Starship{}, nil }
func (s *swApiClientMock) Species(id int) (swapi.Species, error)   { return swapi.Species{}, nil }
func (s *swApiClientMock) Person(id int) (swapi.Person, error)     { return swapi.Person{}, nil }

func (s *swApiClientMock) Planet(id int) (swapi.Planet, error) {
	if s.dropError {
		return swapi.Planet{}, errors.New("some error")
	}
	return swapi.Planet{
		Name:    "Terra",
		Climate: "tropical",
		Terrain: "terra",
		FilmURLs: []string{
			"https://something.com/api/film/1/",
		},
		URL: "https://something.com/api/planet/1/",
	}, nil
}
func (s *swApiClientMock) Film(id int) (swapi.Film, error) {
	if s.dropError {
		return swapi.Film{}, errors.New("some error")
	}
	return swapi.Film{
		Title:      "Filme da terra",
		EpisodeID:  1,
		Created:    "800 quintilhões de anos atras",
		Director:   "Único",
		PlanetURLs: []string{"https://something.com/api/planet/1/"},
		URL:        "https://something.com/api/film/1/",
	}, nil
}

func (c *searchableHttpClientMock) Do(req *http.Request) (*http.Response, error) {
	h := make(http.Header)
	h.Add("Content-Type", "application/json")
	if strings.Contains(req.RequestURI, "error") {
		return &http.Response{
			Status:     "500 Error",
			StatusCode: 500,
			Header:     h,
			Request:    req,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"some error"}`)), //type is io.ReadCloser,
		}, errors.New("some error")
	}

	r := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header:     h,
		Request:    req,
	}
	if strings.Contains(req.RequestURI, "film") {
		r.Body = ioutil.NopCloser(strings.NewReader(`{
				"title": "Filme da terra",
				"episode_id": 1,
				"director": "Único",
				"planets": ["https://something.com/api/planet/1/"],
				"created": "800 quintilhões de anos atras",
				"url": "https://something.com/api/film/1/"
			}`))
	} else {
		r.Body = ioutil.NopCloser(strings.NewReader(`{
				"name": "Terra",
				"climate": "tropical",
				"terrain": "terra",
				"films": [
					"https://something.com/api/film/1/"
				],
				"url": "https://something.com/api/planet/1/"
		}`))
	}
	return r, nil
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

func Test_swApiServiceFinal_Start(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	sm, ctx := NewManagerForTests()
	swapiaccessService := sm.WithSwApiService(NewSwService()).SwApiService()
	tests := []struct {
		name    string
		n       *swApiServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success", //just success because SwApiService.Start(ctx) does nothing with the ctx for now..
			n:       swapiaccessService.(*swApiServiceFinal),
			args:    args{ctx: ctx},
			wantErr: false,
		},
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
	sm, ctx := NewManagerForTests()
	swapiaccessService := sm.WithSwApiService(NewSwService()).SwApiService()
	tests := []struct {
		name    string
		n       *swApiServiceFinal
		args    args
		wantErr bool
	}{
		{
			name:    "success", //just success because SwApiService.Close(ctx) does nothing with the ctx for now..
			n:       swapiaccessService.(*swApiServiceFinal),
			args:    args{ctx: ctx},
			wantErr: false,
		},
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
	sm, ctx := NewManagerForTests()
	tests := []struct {
		name    string
		n       *swApiServiceFinal
		args    args
		wantErr bool
	}{
		{
			name: "success",
			n: sm.WithSwApiService(&swApiServiceFinal{swclient: &swApiClientMock{}, searchableHttpClient: &searchableHttpClientMock{}, online: true}).
				SwApiService().(*swApiServiceFinal),
			args:    args{ctx: ctx},
			wantErr: false,
		},
		{
			name: "error",
			n: sm.WithSwApiService(&swApiServiceFinal{swclient: &swApiClientMock{dropError: true}, searchableHttpClient: &searchableHttpClientMock{}, online: true}).
				SwApiService().(*swApiServiceFinal),
			args:    args{ctx: ctx},
			wantErr: true,
		},
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
	sm, _ := NewManagerForTests()
	swapiaccessService := NewSwService()
	tests := []struct {
		name string
		n    *swApiServiceFinal
		args args
		want services.SwApiService
	}{
		{
			name: "success",
			n:    swapiaccessService.(*swApiServiceFinal),
			args: args{sm: sm},
			want: swapiaccessService,
		},
		{
			name: "nil",
			n:    swapiaccessService.(*swApiServiceFinal),
			args: args{sm: nil},
			want: swapiaccessService,
		},
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
	sm, _ := NewManagerForTests()
	swapiaccessService := sm.WithSwApiService(NewSwService()).SwApiService()
	tests := []struct {
		name string
		n    *swApiServiceFinal
		want services.ServiceManager
	}{
		{
			name: "success",
			n:    swapiaccessService.(*swApiServiceFinal),
			want: sm,
		},
		{
			name: "nil",
			n:    NewSwService().(*swApiServiceFinal),
			want: nil,
		},
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
	sm, _ := NewManagerForTests()
	swapiaccessService := sm.WithSwApiService(NewSwService()).SwApiService()
	tests := []struct {
		name    string
		n       *swApiServiceFinal
		args    args
		want    *models.Planet
		wantErr bool
	}{
		{
			name: "success",
			n:    swapiaccessService.(*swApiServiceFinal),
			args: args{id: 1},
			want: basicPlanet,
		},
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

func planetSuperficialDeepEqual(p1 *models.Planet, p2 *models.Planet) bool {
	return p1.Id == p2.Id && p1.Name == p2.Name && p1.Terrain == p2.Terrain && p1.Climate == p2.Climate
}
