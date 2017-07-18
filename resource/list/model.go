package list

type SEOut struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   []SE   `json:"data"`
}
type SE struct {
	Season  string `json:"season"`
	Episode string `json:"episode"`
}

type ItemListOut struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   []Item `json:"data"`
}

type Item struct {
	ID      string     `json:"id"`
	Name    string     `json:"name"`
	Format  string     `json:"format"`
	Season  string     `json:"season"`
	Episode string     `json:"episode"`
	Size    string     `json:"size"`
	Link    []ItemLink `json:"link"`
}

type ItemLink struct {
	Way     string `json:"way"`
	Address string `json:"address"`
}
