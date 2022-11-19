package swapiaccess

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/marcosArruda/swapi/pkg/messages"
	"github.com/marcosArruda/swapi/pkg/models"
	"github.com/marcosArruda/swapi/pkg/services"
	"github.com/peterhellberg/swapi"
)

type (
	swApiServiceFinal struct {
		sm                   services.ServiceManager
		swclient             *swapi.Client
		online               bool
		searchableHttpClient *http.Client
	}
)

func NewSwService(ctx context.Context) services.SwApiService {
	return &swApiServiceFinal{swclient: swapi.DefaultClient}
}

func (n *swApiServiceFinal) Start(ctx context.Context) error {
	n.sm.LogsService().Info(ctx, "SwApi Service Started Started!")
	return nil
}

func (n *swApiServiceFinal) Close(ctx context.Context) error {
	return nil
}

func (n *swApiServiceFinal) Healthy(ctx context.Context) error {
	if n.online {
		_, err := n.swclient.Planet(1)
		if err != nil {
			n.sm.LogsService().Warn(ctx, messages.SwApiUnavailableError.Error())
			n.online = false
			return messages.SwApiUnavailableError
		}
	}
	return nil
}

func (n *swApiServiceFinal) WithServiceManager(sm services.ServiceManager) services.SwApiService {
	n.sm = sm
	return n
}

func (n *swApiServiceFinal) ServiceManager() services.ServiceManager {
	return n.sm
}

func (n *swApiServiceFinal) GetPlanetById(ctx context.Context, id int) (*models.Planet, error) {
	if n.online {
		p, err := n.swclient.Planet(id)
		if err != nil {
			return nil, messages.SwApiUnavailableError
		}
		return n.ToPersistentPlanet(ctx, &p, id, true)
	}
	return nil, messages.SwApiIsOfflineError
}

func (n *swApiServiceFinal) GetFilmById(ctx context.Context, id int) (*models.Film, error) {
	if n.online {
		f, err := n.swclient.Film(id)
		if err != nil {
			return nil, messages.SwApiUnavailableError
		}
		return n.ToPersistentFilm(ctx, &f, id, true)
	}
	return nil, messages.SwApiIsOfflineError
}

func (n *swApiServiceFinal) SearchPlanetsByName(ctx context.Context, name string) ([]*models.Planet, error) {
	uurl := &url.URL{
		Scheme: "https",
		Host:   "swapi.dev",
	}
	rel, err := url.Parse(fmt.Sprintf("/api/planets/?search=%s", name))

	if err != nil {
		return services.EmptyPlanetSlice, err
	}
	q := rel.Query()
	q.Set("format", "json")
	rel.RawQuery = q.Encode()
	u := uurl.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return services.EmptyPlanetSlice, err
	}
	req.Header.Add("User-Agent", "swapiapp.go")
	req.Close = true
	resp, err := n.searchableHttpClient.Do(req)
	if err != nil {
		return services.EmptyPlanetSlice, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	//TODO: Finish this

	//return req, nil
	return services.EmptyPlanetSlice, nil
}

func (n *swApiServiceFinal) ToPersistentPlanet(ctx context.Context, p *swapi.Planet, id int, expand bool) (*models.Planet, error) {
	pp := &models.Planet{
		Name:     p.Name,
		Climate:  p.Climate,
		Terrain:  p.Terrain,
		FilmURLs: p.FilmURLs,
		Films:    []*models.Film{},
		URL:      p.URL,
	}
	if id < 1 {
		s := strings.Split(p.URL, "/")
		idTmp, err := strconv.Atoi(s[len(s)-2])
		if err != nil {
			return nil, &messages.PlanetError{
				Msg: fmt.Sprintf("Could not discover the planet named '%s' respective ID from the payload", p.Name)}
		}
		pp.Id = idTmp
		return pp, nil
	}
	pp.Id = id
	if expand {
		pp = n.fillFilmsFromPlanet(ctx, pp)
	}
	return pp, nil
}

func (n *swApiServiceFinal) fillFilmsFromPlanet(ctx context.Context, p *models.Planet) *models.Planet {
	if p.FilmURLs != nil && len(p.FilmURLs) > 0 {
		for _, fUrl := range p.FilmURLs {
			s := strings.Split(fUrl, "/")
			idTmp, err := strconv.Atoi(s[len(s)-2])
			if err != nil {
				continue
			}
			ff, err := n.swclient.Film(idTmp)
			if err != nil {
				n.sm.LogsService().Warn(ctx, messages.SwApiUnavailableError.Error())
				continue
			}
			f, err := n.ToPersistentFilm(ctx, &ff, idTmp, false)
			if err != nil {
				n.sm.LogsService().Warn(ctx, err.Error())
				continue
			}
			p.Films = append(p.Films, f)
		}
	}
	return p
}

func (n *swApiServiceFinal) ToPersistentFilm(ctx context.Context, f *swapi.Film, id int, expand bool) (*models.Film, error) {
	ff := &models.Film{
		Title:      f.Title,
		EpisodeID:  f.EpisodeID,
		Director:   f.Director,
		PlanetURLs: f.PlanetURLs,
		Planets:    []*models.Planet{},
		Created:    f.Created,
		URL:        f.URL,
	}
	if id < 1 {
		s := strings.Split(f.URL, "/")
		idTmp, err := strconv.Atoi(s[len(s)-2])
		if err != nil {
			return nil, &messages.PlanetError{
				Msg: fmt.Sprintf("Could not discover the film named '%s' respective ID from the payload", f.Title)}
		}
		ff.Id = idTmp
		return ff, nil
	}
	ff.Id = id
	if expand {
		n.fillPlanetsFromFilm(ctx, ff)
	}
	return ff, nil
}

func (n *swApiServiceFinal) fillPlanetsFromFilm(ctx context.Context, f *models.Film) *models.Film {
	if f.PlanetURLs != nil && len(f.PlanetURLs) > 0 {
		for _, pUrl := range f.PlanetURLs {
			s := strings.Split(pUrl, "/")
			idTmp, err := strconv.Atoi(s[len(s)-2])
			if err != nil {
				continue
			}
			pp, err := n.swclient.Planet(idTmp)
			if err != nil {
				n.sm.LogsService().Warn(ctx, messages.SwApiUnavailableError.Error())
				continue
			}
			p, err := n.ToPersistentPlanet(ctx, &pp, idTmp, false)
			if err != nil {
				n.sm.LogsService().Warn(ctx, err.Error())
				continue
			}
			f.Planets = append(f.Planets, p)
		}
	}
	return f
}

func (n *swApiServiceFinal) PutOnline() services.SwApiService {
	n.online = true
	return n
}

func (n *swApiServiceFinal) PutOffline() services.SwApiService {
	n.online = false
	return n
}

func (n *swApiServiceFinal) IsOnline() bool {
	return n.online
}
