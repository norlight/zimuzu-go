package top

type TopOut struct {
	Status int        `json:"status"`
	Info   string     `json:"info"`
	Data   []Resource `json:"data"`
}

type Resource struct {
	ID          string `json:"id"`
	Cnname      string `json:"cnname"`
	Channel     string `json:"channel"`
	Area        string `json:"area"`
	Category    string `json:"category"`
	PublishYear string `json:"publish_year"`
	PlayStatus  string `json:"play_status"`
	Poster      string `json:"poster"`
	PosterB     string `json:"poster_b"`
	PosterM     string `json:"poster_m"`
	PosterS     string `json:"poster_s"`
}
