package messages

import (
	"errors"
)

var (
	SwApiUnavailableError = errors.New("something went wrong accessing swapi.dev. I will change the sistem to 'offline' mode for now")
	SwApiIsOfflineError   = errors.New("SwApi is in 'offline' mode now.")
	NoPlanetFound         = errors.New("No Planet found")
)

type (
	PlanetError struct {
		Msg        string
		PlanetName string
		PlanetId   int
	}

	FilmError struct {
		Msg       string
		FilmTitle string
		FilmId    int
	}
)

func (p *PlanetError) Error() string {
	return p.Msg
}

func (f *FilmError) Error() string {
	return f.Msg
}
