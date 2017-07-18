package resource

//Subject 资源，剧、电影等
type Subject struct {
	ID         string `json:"id"`
	CNName     string `json:"cnname"`
	ENName     string `json:"enname"`
	Remark     string `json:"remark"`
	Area       string `json:"area"`
	Format     string `json:"format"`
	Category   string `json:"category"`
	Poster     string `json:"poster"`
	Channel    string `json:"channel"`
	Lang       string `json:"lang"`
	PlayStatus string `json:"play_status"`
	Rank       string `json:"rank"`
	Score      string `json:"score"`
	Views      string `json:"views"`
	Itemupdate string `json:"itemupdate"`
	PosterA    string `json:"poster_a"`
	PosterB    string `json:"poster_b"`
	PosterM    string `json:"poster_m"`
	PosterS    string `json:"poster_s"`
}

//List 资源列表
type List struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   struct {
		Count string    `json:"count"`
		List  []Subject `json:"list"`
	} `json:"data"`
}

//Info 资源详情
type Info struct {
	Status int     `json:"status"`
	Info   string  `json:"info"`
	Data   Subject `json:"data"`
}

//SeasonEpisodeInfo 资源季度信息列表
type SeasonEpisodeInfo struct {
	Status int             `json:"status"`
	Info   string          `json:"info"`
	Data   []SeasonEpisode `json:"data"`
}

//SeasonEpisode 资源季度信息
type SeasonEpisode struct {
	Season  string `json:"season"`
	Episode string `json:"episode"`
}

//Item 单集下载信息
type Item struct {
	ID      string     `json:"id"`
	Name    string     `json:"name"`
	Format  string     `json:"format"`
	Season  string     `json:"season"`
	Episode string     `json:"episode"`
	Size    string     `json:"size"`
	Link    []ItemLink `json:"link"`
}

//ItemLink 单集下载地址
type ItemLink struct {
	Way     string `json:"way"`
	Address string `json:"address"`
}

//ItemList 下载列表
type ItemList struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   []struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Format  string `json:"format"`
		Season  string `json:"season"`
		Episode string `json:"episode"`
		Size    string `json:"size"`
		Link    []struct {
			Way     string `json:"way"`
			Address string `json:"address"`
		} `json:"link"`
	} `json:"data"`
}
