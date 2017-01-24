package search

import (
	"net/url"
	"reflect"
	"strings"
)

// Query holds search fields that can be sent to the search service. It encodes
// itself into a URL query string when the search request is made.
type Query struct {
	// Mandatory fields
	DeviceType string `param:"device_type"`
	Language   string `param:"lang"`
	Site       string `param:"site"`

	// Optional fields
	BrandID  string   `param:"brand_id"`
	Episode  string   `param:"episode"`
	PageSize string   `param:"page_size"`
	Season   string   `param:"season"`
	SeasonID string   `param:"season_id"`
	Type     string   `param:"type"`
	VideoIDs []string `param:"video_ids"`

	// Sorting
	SortBy string `param:"sort_by"`
	Order  string `param:"order"`
}

// rawURLQuery encodes the query into a raw URL query string.
func (q *Query) rawURLQuery() string {
	t := reflect.TypeOf(*q)
	s := reflect.ValueOf(*q)

	values := &url.Values{}

	for i := 0; i < t.NumField(); i++ {
		if param, ok := t.Field(i).Tag.Lookup("param"); ok {
			var strVal string

			switch val := s.Field(i).Interface().(type) {
			case string:
				strVal = val
			case []string:
				strVal = strings.Join(val, ",")
			}

			if strVal != "" {
				values.Set(param, strVal)
			}
		}
	}
	return values.Encode()
}
