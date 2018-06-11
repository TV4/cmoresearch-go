package cmoresearch

import (
	"mime"
	"net/http"
	"net/url"
)

var (
	defaultBaseURL = &url.URL{
		Scheme: "https",
		Host:   "cmore-search.b17g.services",
	}
)

// Client is a client for the search service.
type Client struct {
	appName    string
	baseURL    *url.URL
	httpClient *http.Client
	debugLogf  func(string, ...interface{})
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

	if c.debugLogf != nil {
		if c.httpClient.Transport == nil {
			c.httpClient.Transport = http.DefaultTransport
		}
		underlyingTransport := c.httpClient.Transport
		c.httpClient.Transport = roundTripFunc(
			func(r *http.Request) (*http.Response, error) {
				c.debugLogf("%s %s", r.Method, r.URL.String())
				return underlyingTransport.RoundTrip(r)
			},
		)
	} else {
		c.debugLogf = func(string, ...interface{}) {}
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

// SetDebugLogf is an option to set a debug logger when creating a new client.
// If set the client will log requests using this logger, for debug uses.
func SetDebugLogf(logf func(format string, v ...interface{})) func(*Client) {
	return func(c *Client) {
		c.debugLogf = logf
	}
}

// SetAppName is an option to set the value used in the client parameter sent to
// cmore-search.
func SetAppName(appName string) func(*Client) {
	return func(c *Client) {
		c.appName = appName
	}
}

func isJSONResponse(resp *http.Response) bool {
	ct := resp.Header.Get("Content-Type")
	mt, _, _ := mime.ParseMediaType(ct)
	return mt == "application/json"
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (rt roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return rt(r)
}
