package resource

import (
	"net/http"
	"net/url"

	"github.com/norlight/zimuzu-go/client"
)

type Resource struct {
	Client *client.Client
}

func New(c *client.Client) Resource {
	return Resource{c}
}

func (r *Resource) FetchList(channel, area, sort, year, category, limit, page string) (resp *http.Response, err error) {
	p := "/resource/fetchlist"

	q := url.Values{}
	q.Set("channel", channel)
	q.Set("area", area)
	q.Set("sort", sort)
	q.Set("year", year)
	q.Set("category", category)
	q.Set("limit", limit)
	q.Set("page", page)

	resp, err = r.Client.Get(p, q.Encode())
	return
}

func (r *Resource) Top(channel string, limit string) (resp *http.Response, err error) {
	p := "/resource/top"

	q := url.Values{}
	q.Set("channel", channel)
	q.Set("limit", limit)

	resp, err = r.Client.Get(p, q.Encode())
	return
}

func (r *Resource) List(id string) (resp *http.Response, err error) {
	p := "/resource/season_episode"
	q := url.Values{}
	q.Set("id", id)

	resp, err = r.Client.Get(p, q.Encode())
	return
}

func (r *Resource) ItemList(id, season, episode, file string) (resp *http.Response, err error) {
	p := "/resource/itemlist_web"
	q := url.Values{}
	q.Set("id", id)
	q.Set("season", season)
	q.Set("episode", episode)
	q.Set("file", file)

	resp, err = r.Client.Get(p, q.Encode())
	return
}
