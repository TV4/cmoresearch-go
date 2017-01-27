// Package search implements a client for Bonnier Broadcasting's search service.
package search

import (
	"mime"
	"net/http"
	"net/url"
)

// Search is a client for the search service.
type Search struct {
	baseURL    *url.URL
	httpClient *http.Client
	logf       func(string, ...interface{})
}

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

// New returns a new search client.
func New(options ...func(*Search) error) (*Search, error) {
	bu, err := url.Parse("https://search.b17g.services/")
	if err != nil {
		return nil, err
	}

	s := &Search{baseURL: bu}

	for _, o := range options {
		if err := o(s); err != nil {
			return nil, err
		}
	}

	if s.httpClient == nil {
		dup := *http.DefaultClient
		s.httpClient = &dup
	}

	if s.logf == nil {
		s.logf = func(string, ...interface{}) {}
	}

	return s, nil
}

// SetBaseURL is an option to set a custom URL to the search service when
// creating a new Search instance.
func SetBaseURL(rawurl string) func(*Search) error {
	return func(s *Search) error {
		bu, err := url.Parse(rawurl)
		if err != nil {
			return err
		}
		s.baseURL = bu
		return nil
	}
}

// SetHTTPClient is an option to set a custom HTTP client when creating a new
// Search instance.
func SetHTTPClient(c *http.Client) func(*Search) error {
	return func(s *Search) error {
		s.httpClient = c
		return nil
	}
}

// SetLogf is an option to configure a logf (Printf function for logging) when
// creating a new Search instance.
func SetLogf(logf func(string, ...interface{})) func(*Search) error {
	return func(s *Search) error {
		s.logf = logf
		return nil
	}
}

func isJSONResponse(resp *http.Response) bool {
	ct := resp.Header.Get("Content-Type")
	mt, _, _ := mime.ParseMediaType(ct)
	return mt == "application/json"
}
