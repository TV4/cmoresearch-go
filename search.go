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

// Search performs a search and returns the response. An error is returned if
// there is an error while setting up or sending the request, but also if the
// response status is not HTTP 200 OK or the response content is not JSON.
func (c *Client) Search(ctx context.Context, query url.Values, options ...func(r *http.Request)) (Response, error) {
	rel, err := url.Parse(path.Join(c.baseURL.Path, "/search"))
	if err != nil {
		return Response{}, err
	}

	u := c.baseURL.ResolveReference(rel)
	u.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Response{}, err
	}

	req = req.WithContext(ctx)

	for _, o := range options {
		o(req)
	}

	c.logf("GET %s", u)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Response{}, err
	}

	meta := Meta{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
	}

	defer func() {
		io.CopyN(ioutil.Discard, resp.Body, 64)
		resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		if !isJSONResponse(resp) {
			return Response{Meta: meta}, fmt.Errorf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		}
		var ae APIError
		if err := json.NewDecoder(resp.Body).Decode(&ae); err != nil {
			return Response{Meta: meta}, fmt.Errorf("%d %s; JSON response body malformed (%v)", resp.StatusCode, http.StatusText(resp.StatusCode), err)
		}
		return Response{Meta: meta}, &ae
	}

	if !isJSONResponse(resp) {
		return Response{Meta: meta}, errors.New("Content-Type not JSON")
	}

	response, err := makeResponse(resp)
	if err != nil {
		return Response{Meta: meta}, err
	}
	return response, nil
}

// SetRequestID is an option for Search to set the X-Request-Id header on the
// search request.
func SetRequestID(requestID string) func(*http.Request) {
	return func(r *http.Request) {
		r.Header.Set("X-Request-Id", requestID)
	}
}

func makeResponse(resp *http.Response) (Response, error) {
	meta := Meta{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
	}

	var v struct {
		TotalHits int               `json:"total_hits"`
		Hits      []json.RawMessage `json:"assets"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return Response{Meta: meta}, err
	}

	response := Response{TotalHits: v.TotalHits}

	for _, h := range v.Hits {
		var t struct {
			Type string
		}

		if json.Unmarshal(h, &t) == nil {
			switch t.Type {
			case "series":
				var series Series
				if json.Unmarshal(h, &series) == nil {
					response.Hits = append(response.Hits, &series)
				}
			default:
				var asset Asset
				if json.Unmarshal(h, &asset) == nil {
					response.Hits = append(response.Hits, &asset)
				}
			}
		}
	}

	response.Meta = meta
	return response, nil
}
