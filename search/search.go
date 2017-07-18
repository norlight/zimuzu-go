package search

import (
	"net/http"
	"net/url"

	"github.com/norlight/zimuzu-go/client"
)

type Search struct {
	Client *client.Client
}

func New(c *client.Client) Search {
	return Search{c}
}

func (s *Search) Search(k string) (resp *http.Response, err error) {
	p := "/search"
	q := url.Values{}
	q.Set("k", k)

	resp, err = s.Client.Get(p, q.Encode())
	return
}
