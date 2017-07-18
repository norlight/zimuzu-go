package client

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	Host = "http://api.ousns.net"
)

type Client struct {
	cid       string
	accesskey string
	platform  int
	client    *http.Client
}

func New(cid string, key string) (c Client) {
	c = Client{
		cid:       cid,
		accesskey: key,
		platform:  1,
		client:    &http.Client{},
	}
	return
}

func (c *Client) Get(p string, q string) (resp *http.Response, err error) {
	client := c.client
	u, err := url.Parse(Host)
	if err != nil {
		return nil, err
	}

	u.Path = p

	query, err := url.ParseQuery(q)
	if err != nil {
		return nil, err
	}
	query.Set("cid", c.cid)
	query.Set("accesskey", c.encryptKey())
	query.Set("timestamp", fmt.Sprintf("%d", time.Now().Unix()))
	query.Set("client", fmt.Sprintf("%d", c.platform))

	u.RawQuery = query.Encode()

	//log.Println(u.String())
	r, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err = client.Do(r)
	return
}

func (c *Client) encryptKey() string {
	str := fmt.Sprintf("%s$$%s&&%d", c.cid, c.accesskey, time.Now().Unix())
	//m := md5.Sum([]byte(str))
	m := md5.New()
	io.WriteString(m, str)
	md5 := m.Sum(nil)
	return fmt.Sprintf("%x", md5)
}
