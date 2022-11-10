package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/marcosArruda/swapi/pkg/messages"
	"github.com/peterhellberg/swapi"
)

type Film struct {
	Id         int
	Title      string   `json:"title"`
	EpisodeID  int      `json:"episode_id"`
	Director   string   `json:"director"`
	PlanetURLs []string `json:"planets"`
	Created    string   `json:"created"`
	URL        string   `json:"url"`
}

func ToPersistentFilm(f *swapi.Film, id int) (*Film, error) {
	ff := &Film{
		Title:      f.Title,
		EpisodeID:  f.EpisodeID,
		Director:   f.Director,
		PlanetURLs: f.PlanetURLs,
		Created:    f.Created,
		URL:        f.URL,
	}
	if id == 0 {
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
	return ff, nil
}
