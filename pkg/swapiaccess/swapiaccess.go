package swapiaccess

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

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
	return &swApiServiceFinal{swclient: swapi.DefaultClient, searchableHttpClient: http.DefaultClient}
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
	r := &models.SwApiPlanetsByNameResult{Next: fmt.Sprintf("https://swapi.dev/api/planets/?search=%s", name)}
	ps := []*models.Planet{}
	for r.Next != "" {
		rel, err := url.Parse(r.Next)
		if err != nil {
			return services.EmptyPlanetSlice, err
		}
		q := rel.Query()
		q.Set("format", "json")
		rel.RawQuery = q.Encode()
		req, err := http.NewRequest("GET", rel.String(), nil)
		if err != nil {
			return services.EmptyPlanetSlice, err
		}
		req.Header.Add("User-Agent", "swapiapp.go")
		req.Close = true
		n.sm.LogsService().Info(ctx, fmt.Sprintf("Calling r.Next -> %s", r.Next))
		resp, err := n.searchableHttpClient.Do(req)
		if err != nil {
			return services.EmptyPlanetSlice, err
		}
		defer resp.Body.Close()
		n.sm.LogsService().Info(ctx, fmt.Sprintf("Got %d total results from r.Next (%s)", r.Count, r.Next))
		r.Next = "" //zeroing the Next to avoid infinite loop
		if err = json.NewDecoder(resp.Body).Decode(r); err != nil {
			return services.EmptyPlanetSlice, err
		}
		n.sm.LogsService().Info(ctx, fmt.Sprintf("Decoding for r.Next (%s) worked!", r.Next))
		for _, rr := range r.Results {
			n.sm.LogsService().Info(ctx, fmt.Sprintf("Converting '%s' to persistent", rr.Name))
			pp, err := n.ToPersistentPlanet(ctx, rr, 0, true)
			if err != nil {
				return services.EmptyPlanetSlice, err
			}
			ps = append(ps, pp)
		}

	}
	for _, v := range ps {
		fmt.Printf("planet {id: %d, name: %s}", v.Id, v.Name)
		fmt.Println("")
	}
	return ps, nil
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
		n.sm.LogsService().Info(ctx, fmt.Sprintf("id for planet '%s' is %d and it have %d films", p.Name, idTmp, len(p.FilmURLs)))
		pp.Id = idTmp
	} else {
		pp.Id = id
	}

	if expand {
		pp = n.fillFilmsFromPlanet(ctx, pp)
	}
	return pp, nil
}

func (n *swApiServiceFinal) fillFilmsFromPlanet(ctx context.Context, p *models.Planet) *models.Planet {
	var wg sync.WaitGroup
	if p.FilmURLs != nil && len(p.FilmURLs) > 0 {
		n.sm.LogsService().Info(ctx, fmt.Sprintf("Planet '%s' have %d films, expanding them ..", p.Name, len(p.FilmURLs)))
		for _, fUrl := range p.FilmURLs {
			s := strings.Split(fUrl, "/")
			idTmp, err := strconv.Atoi(s[len(s)-2])
			if err != nil {
				continue
			}

			wg.Add(1)
			go func() {
				ff, err := n.swclient.Film(idTmp)
				if err != nil {
					n.sm.LogsService().Warn(ctx, messages.SwApiUnavailableError.Error())
				}
				f, err := n.ToPersistentFilm(ctx, &ff, idTmp, false)
				if err != nil {
					n.sm.LogsService().Warn(ctx, err.Error())
				}
				p.Films = append(p.Films, f)
				wg.Done()
			}()
		}
	}
	wg.Wait()
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
	var wg sync.WaitGroup
	if f.PlanetURLs != nil && len(f.PlanetURLs) > 0 {
		for _, pUrl := range f.PlanetURLs {
			s := strings.Split(pUrl, "/")
			idTmp, err := strconv.Atoi(s[len(s)-2])
			if err != nil {
				continue
			}
			wg.Add(1)
			go func() {
				pp, err := n.swclient.Planet(idTmp)
				if err != nil {
					n.sm.LogsService().Warn(ctx, messages.SwApiUnavailableError.Error())
				}
				p, err := n.ToPersistentPlanet(ctx, &pp, idTmp, false)
				if err != nil {
					n.sm.LogsService().Warn(ctx, err.Error())
				}
				f.Planets = append(f.Planets, p)
				wg.Done()
			}()
		}
	}
	wg.Wait()
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
