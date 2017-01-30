package search

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// Do performs a search and returns the result.
func (c *Client) Do(ctx context.Context, query url.Values, options ...func(r *http.Request)) (*Result, error) {
	rel, err := url.Parse(path.Join(c.baseURL.Path, "/search"))
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)
	u.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	for _, o := range options {
		o(req)
	}

	c.logf("GET %s", u)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		io.CopyN(ioutil.Discard, resp.Body, 64)
		resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		if !isJSONResponse(resp) {
			return nil, fmt.Errorf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		}
		var ae APIError
		if err := json.NewDecoder(resp.Body).Decode(&ae); err != nil {
			return nil, fmt.Errorf("%d %s; JSON response body malformed (%v)", resp.StatusCode, http.StatusText(resp.StatusCode), err)
		}
		return nil, &ae
	}

	if !isJSONResponse(resp) {
		return nil, errors.New("Content-Type not JSON")
	}

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SetRequestID is an option for Do to set the X-Request-Id header on the
// search request.
func SetRequestID(requestID string) func(*http.Request) {
	return func(r *http.Request) {
		r.Header.Set("X-Request-Id", requestID)
	}
}
