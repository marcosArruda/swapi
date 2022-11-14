package models

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/marcosArruda/swapi/pkg/messages"
	"github.com/peterhellberg/swapi"
)

type (
	Planet struct {
		Id       int      `json:"id"`
		Name     string   `json:"name"`
		Climate  string   `json:"climate"`
		Terrain  string   `json:"terrain"`
		FilmURLs []string `json:"film_urls"`
		Films    []*Film  `json:"films"`
		URL      string   `json:"url"`
	}
)

func ToPersistentPlanet(ctx context.Context, p *swapi.Planet, id int) (*Planet, error) {
	pp := &Planet{
		Name:     p.Name,
		Climate:  p.Climate,
		Terrain:  p.Terrain,
		FilmURLs: p.FilmURLs,
		URL:      p.URL,
	}
	if id == 0 {
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
	return pp, nil
}
