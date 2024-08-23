package v2

import (
	"Lbt-Management/pkg/qbittorrent"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	httpClient *http.Client
	host       string
	api        string
}

func NewClient(api string) *Client {
	return &Client{
		httpClient: &http.Client{},
		host:       qbittorrent.V2BaseUrl,
		api:        api,
	}
}

func (c *Client) NewRequest(method string, params map[string]string) (*http.Request, error) {
	fullURL := c.host + c.api

	var req *http.Request
	var err error

	if method == http.MethodGet {
		u, err := url.Parse(fullURL)
		if err != nil {
			return nil, err
		}

		q := u.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		u.RawQuery = q.Encode()

		req, err = http.NewRequest(http.MethodGet, u.String(), nil)
	} else {
		form := url.Values{}
		for key, value := range params {
			form.Set(key, value)
		}

		req, err = http.NewRequest(method, fullURL, strings.NewReader(form.Encode()))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}
