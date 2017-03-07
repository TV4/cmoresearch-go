// Package search implements a client for Bonnier Broadcasting's search service.
package search

import (
	"mime"
	"net/http"
	"net/url"
)

var (
	defaultBaseURL = &url.URL{
		Scheme: "https",
		Host:   "search.b17g.services",
	}
)

// Client is a client for the search service.
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

// NewClient returns a new search client.
func NewClient(options ...func(*Client)) *Client {
	c := &Client{baseURL: defaultBaseURL}

	for _, o := range options {
		o(c)
	}

	if c.httpClient == nil {
		dup := *http.DefaultClient
		c.httpClient = &dup
	}

	return c
}

// SetBaseURL is an option to set a custom URL to the search service when
// creating a new Search instance.
func SetBaseURL(rawurl string) func(*Client) {
	return func(c *Client) {
		if bu, err := url.Parse(rawurl); err == nil {
			c.baseURL = bu
			return
		}
		c.baseURL = nil
	}
}

// SetHTTPClient is an option to set a custom HTTP client when creating a new
// Search instance.
func SetHTTPClient(hc *http.Client) func(*Client) {
	return func(c *Client) {
		c.httpClient = hc
	}
}

func isJSONResponse(resp *http.Response) bool {
	ct := resp.Header.Get("Content-Type")
	mt, _, _ := mime.ParseMediaType(ct)
	return mt == "application/json"
}
