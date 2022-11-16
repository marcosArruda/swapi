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
)
