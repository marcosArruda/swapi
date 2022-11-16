package models

type Film struct {
	Id         int
	Title      string    `json:"title"`
	EpisodeID  int       `json:"episode_id"`
	Director   string    `json:"director"`
	PlanetURLs []string  `json:"-"`
	Planets    []*Planet `json:"planets"`
	Created    string    `json:"created"`
	URL        string    `json:"url"`
}
