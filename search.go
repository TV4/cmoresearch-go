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

// Result represents the result as received from the search service.
type Result struct {
	TotalHits int      `json:"total_hits"`
	Assets    []*Asset `json:"assets"`
}

// Asset is a subset of an asset returned by the search service. Its structure
// matches the response format of the search service.
type Asset struct {
	Brand                 Brand     `json:"brand"`
	DescriptionExtendedDa string    `json:"description_extended_da"`
	DescriptionExtendedFi string    `json:"description_extended_fi"`
	DescriptionExtendedNb string    `json:"description_extended_nb"`
	DescriptionExtendedSv string    `json:"description_extended_sv"`
	DescriptionMediumDa   string    `json:"description_medium_da"`
	DescriptionMediumFi   string    `json:"description_medium_fi"`
	DescriptionMediumNb   string    `json:"description_medium_nb"`
	DescriptionMediumSv   string    `json:"description_medium_sv"`
	DescriptionShortDa    string    `json:"description_short_da"`
	DescriptionShortFi    string    `json:"description_short_fi"`
	DescriptionShortNb    string    `json:"description_short_nb"`
	DescriptionShortSv    string    `json:"description_short_sv"`
	EpisodeNumber         int       `json:"episode_number"`
	ID                    string    `json:"video_id"`
	Landscape             Landscape `json:"landscape"`
	Season                Season    `json:"season"`
	Seasons               []int     `json:"seasons"`
	TitleDa               string    `json:"title_da"`
	TitleFi               string    `json:"title_fi"`
	TitleNb               string    `json:"title_nb"`
	TitleSv               string    `json:"title_sv"`
	Type                  string    `json:"type"`
}

// Brand is the brand part of an Asset
type Brand struct {
	ID      string `json:"id"`
	TitleDa string `json:"title_da"`
	TitleFi string `json:"title_fi"`
	TitleNb string `json:"title_nb"`
	TitleSv string `json:"title_sv"`
}

// Landscape is the landscape image of an Asset
type Landscape struct {
	URL string `json:"url"`
}

// Season is the season part of an Asset
type Season struct {
	ID     string `json:"id"`
	Number int    `json:"season_number"`
}

// Search performs a search and returns the result.
func (c *Client) Search(ctx context.Context, query url.Values, options ...func(r *http.Request)) (*Result, error) {
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

// SetRequestID is an option for Search to set the X-Request-Id header on the
// search request.
func SetRequestID(requestID string) func(*http.Request) {
	return func(r *http.Request) {
		r.Header.Set("X-Request-Id", requestID)
	}
}
