package models

type (
	Planet struct {
		Id       int      `json:"id"`
		Name     string   `json:"name"`
		Climate  string   `json:"climate"`
		Terrain  string   `json:"terrain"`
		FilmURLs []string `json:"-"`
		Films    []*Film  `json:"films"`
		URL      string   `json:"url"`
	}

	Film struct {
		Id         int
		Title      string    `json:"title"`
		EpisodeID  int       `json:"episode_id"`
		Director   string    `json:"director"`
		PlanetURLs []string  `json:"-"`
		Planets    []*Planet `json:"planets"`
		Created    string    `json:"created"`
		URL        string    `json:"url"`
	}

	FilmPlanet struct {
		FId int
		PId int
	}

	PlanetByName struct {
		Name string `json:"name"`
	}

	PlanetByExactName struct {
		Name string `json:"name"`
	}
)
