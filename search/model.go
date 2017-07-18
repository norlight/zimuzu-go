package search

type SearchOut struct {
	Status int      `json:"status"`
	Info   string   `json:"info"`
	Data   ItemList `json:"data"`
}

type ItemList struct {
	Count int    `json:"count"`
	List  []Item `json:"list"`
}
type Item struct {
	Itemid  string `json:"itemid"`
	Title   string `json:"title"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Pubtime string `json:"pubtime"`
	Uptime  string `json:"uptime"`
}
