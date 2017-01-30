// Package search implements a client for Bonnier Broadcasting's search service.
package search

import (
	"mime"
	"net/http"
	"net/url"
)

// Client is a client for the search service.
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	logf       func(string, ...interface{})
}

// New returns a new search client.
func New(options ...func(*Client) error) (*Client, error) {
	bu, err := url.Parse("https://search.b17g.services/")
	if err != nil {
		return nil, err
	}

	c := &Client{baseURL: bu}

	for _, o := range options {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	if c.httpClient == nil {
		dup := *http.DefaultClient
		c.httpClient = &dup
	}

	if c.logf == nil {
		c.logf = func(string, ...interface{}) {}
	}

	return c, nil
}

// SetBaseURL is an option to set a custom URL to the search service when
// creating a new Search instance.
func SetBaseURL(rawurl string) func(*Client) error {
	return func(c *Client) error {
		bu, err := url.Parse(rawurl)
		if err != nil {
			return err
		}
		c.baseURL = bu
		return nil
	}
}

// SetHTTPClient is an option to set a custom HTTP client when creating a new
// Search instance.
func SetHTTPClient(hc *http.Client) func(*Client) error {
	return func(c *Client) error {
		c.httpClient = hc
		return nil
	}
}

// SetLogf is an option to configure a logf (Printf function for logging) when
// creating a new Search instance.
func SetLogf(logf func(string, ...interface{})) func(*Client) error {
	return func(c *Client) error {
		c.logf = logf
		return nil
	}
}

func isJSONResponse(resp *http.Response) bool {
	ct := resp.Header.Get("Content-Type")
	mt, _, _ := mime.ParseMediaType(ct)
	return mt == "application/json"
}
