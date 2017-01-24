package search

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type mockTransport func(*http.Request) (*http.Response, error)

func (mt mockTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return mt(r)
}

func TestIsJSONResponse(t *testing.T) {
	for n, tt := range []struct {
		contentType string
		isJSON      bool
	}{
		{"application/json; charset=utf-8", true},
		{"application/json; charset=iso-8859-1", true},
		{"application/json", true},
		{"text/plain", false},
		{"randomnoiseapplication/jsonrandomnoise", false},
	} {
		resp := &http.Response{Header: make(http.Header)}
		resp.Header.Add("Content-Type", tt.contentType)

		if got, want := isJSONResponse(resp), tt.isJSON; got != want {
			t.Errorf("[%d] %q -> got %t, want %t", n, tt.contentType, got, want)
		}
	}
}

func TestDo(t *testing.T) {
	t.Run("ResponseParsing", func(t *testing.T) {
		var mockT mockTransport = func(r *http.Request) (*http.Response, error) {
			const mockResponseBody = `
				{
					"assets": [
						{
							"brand": {
								"id": "34515",
								"title_da": "Danish title",
								"title_fi": "Finnish title",
								"title_nb": "Norwegian title",
								"title_sv": "Swedish title"
							},
							"description_extended_da": "Danish description (extended)",
							"description_extended_fi": "Finnish desription (extended)",
							"description_extended_nb": "Norwegian description (extended)",
							"description_extended_sv": "Swedish description (extended)",
							"description_medium_da": "Danish description (medium)",
							"description_medium_fi": "Finnish desription (medium)",
							"description_medium_nb": "Norwegian description (medium)",
							"description_medium_sv": "Swedish description (medium)",
							"description_short_da": "Danish description (short)",
							"description_short_fi": "Finnish desription (short)",
							"description_short_nb": "Norwegian description (short)",
							"description_short_sv": "Swedish description (short)",
							"episode_number": 9,
							"landscape": {
								"url": "http://example.com/image.jpg"
							},
							"season": {
								"id": "season-id",
								"season_number": 3
							},
							"seasons": [1, 2, 3, 4, 5],
							"title_da": "Danish asset title",
							"title_fi": "Finnish asset title",
							"title_nb": "Norwegian asset title",
							"title_sv": "Swedish asset title",
							"type": "episode",
							"video_id": "2222333"
						}
					],
					"total_hits": 1
				}
			`
			resp := &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader(mockResponseBody)),
				Header:     make(http.Header),
				StatusCode: http.StatusOK,
			}
			resp.Header.Add("Content-Type", "application/json; charset=utf-8")
			return resp, nil
		}

		hc := &http.Client{Transport: mockT}

		s, err := New("/", hc, nil)
		if err != nil {
			t.Fatalf("New: unexpected error: %v", err)
		}

		res, err := s.Do(&Query{})
		if err != nil {
			t.Fatalf("Search: unexpected error: %v", err)
		}

		if got, want := len(res.Assets), 1; got != want {
			t.Fatalf("len(res.Assets) = %d, want %d", got, want)
		}

		if got, want := res.Assets[0].Brand.ID, "34515"; got != want {
			t.Errorf("res.Assets[0].Brand.ID = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].Brand.TitleDa, "Danish title"; got != want {
			t.Errorf("res.Assets[0].Brand.TitleDa = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].Brand.TitleFi, "Finnish title"; got != want {
			t.Errorf("res.Assets[0].Brand.TitleFi = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].Brand.TitleNb, "Norwegian title"; got != want {
			t.Errorf("res.Assets[0].Brand.TitleNb = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].Brand.TitleSv, "Swedish title"; got != want {
			t.Errorf("res.Assets[0].Brand.TitleSv = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].DescriptionExtendedDa, "Danish description (extended)"; got != want {
			t.Errorf("res.Assets[0].DescriptionExtendedDa = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].DescriptionExtendedFi, "Finnish desription (extended)"; got != want {
			t.Errorf("res.Assets[0].DescriptionExtendedFi = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].DescriptionExtendedNb, "Norwegian description (extended)"; got != want {
			t.Errorf("res.Assets[0].DescriptionExtendedNb = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].DescriptionExtendedSv, "Swedish description (extended)"; got != want {
			t.Errorf("res.Assets[0].DescriptionExtendedSv = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].DescriptionMediumDa, "Danish description (medium)"; got != want {
			t.Errorf("res.Assets[0].DescriptionMediumDa = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].DescriptionMediumFi, "Finnish desription (medium)"; got != want {
			t.Errorf("res.Assets[0].DescriptionMediumFi = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].DescriptionMediumNb, "Norwegian description (medium)"; got != want {
			t.Errorf("res.Assets[0].DescriptionMediumNb = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].DescriptionMediumSv, "Swedish description (medium)"; got != want {
			t.Errorf("res.Assets[0].DescriptionMediumSv = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].DescriptionShortDa, "Danish description (short)"; got != want {
			t.Errorf("res.Assets[0].DescriptionShortDa = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].DescriptionShortFi, "Finnish desription (short)"; got != want {
			t.Errorf("res.Assets[0].DescriptionShortFi = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].DescriptionShortNb, "Norwegian description (short)"; got != want {
			t.Errorf("res.Assets[0].DescriptionShortNb = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].DescriptionShortSv, "Swedish description (short)"; got != want {
			t.Errorf("res.Assets[0].DescriptionShortSv = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].EpisodeNumber, 9; got != want {
			t.Errorf("res.Assets[0].EpisodeNumber = %d, want %d", got, want)
		}

		if got, want := res.Assets[0].ID, "2222333"; got != want {
			t.Errorf("res.Assets[0].ID = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].Landscape.URL, "http://example.com/image.jpg"; got != want {
			t.Errorf("res.Assets[0].Landscape.URL = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].Season.ID, "season-id"; got != want {
			t.Errorf("res.Assets[0].Season.ID = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].Season.Number, 3; got != want {
			t.Errorf("res.Assets[0].Season.Number = %d, want %d", got, want)
		}

		if got, want := len(res.Assets[0].Seasons), 5; got != want {
			t.Errorf("len(res.Assets[0].Seasons) = %d, want %d", got, want)
		} else {
			for n := range res.Assets[0].Seasons {
				if got, want := res.Assets[0].Seasons[n], n+1; got != want {
					t.Errorf("res.Assets[0].Seasons[%d] = %d, want %d", n, got, want)
				}
			}
		}

		if got, want := res.Assets[0].TitleDa, "Danish asset title"; got != want {
			t.Errorf("res.Assets[0].TitleDa = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].TitleFi, "Finnish asset title"; got != want {
			t.Errorf("res.Assets[0].TitleFi = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].TitleNb, "Norwegian asset title"; got != want {
			t.Errorf("res.Assets[0].TitleNb = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].TitleSv, "Swedish asset title"; got != want {
			t.Errorf("res.Assets[0].TitleSv = %q, want %q", got, want)
		}

		if got, want := res.Assets[0].Type, "episode"; got != want {
			t.Errorf("res.Assets[0].Type = %q, want %q", got, want)
		}

		if got, want := res.TotalHits, 1; got != want {
			t.Errorf("res.TotalHits = %d, want %d", got, want)
		}
	})

	t.Run("NonAPIErrors", func(t *testing.T) {
		for n, tt := range []struct {
			status       int
			contentType  string
			errorMessage string
		}{
			{http.StatusBadRequest, "text/plain; charset=utf-8", "400 Bad Request"},
			{http.StatusInternalServerError, "text/plain; charset=utf-8", "500 Internal Server Error"},
			{http.StatusOK, "text/hieroglyphs; charset=ucs-2", "Content-Type not JSON"},
		} {
			var mockT mockTransport = func(r *http.Request) (*http.Response, error) {
				resp := &http.Response{
					Body:       ioutil.NopCloser(strings.NewReader("all is lost!")),
					Header:     make(http.Header),
					StatusCode: tt.status,
				}
				resp.Header.Add("Content-Type", tt.contentType)
				return resp, nil
			}

			hc := &http.Client{Transport: mockT}

			s, err := New("/", hc, nil)
			if err != nil {
				t.Fatalf("[%d] New: unexpected error: %v", n, err)
			}

			_, err = s.Do(&Query{})

			if err == nil {
				t.Fatalf("[%d], Search: got nil, want err", n)
			}

			if got, want := err.Error(), tt.errorMessage; got != want {
				t.Errorf("[%d] err.Error() = %q, want %q", n, got, want)
			}
		}
	})

	t.Run("APIError", func(t *testing.T) {
		var mockT mockTransport = func(r *http.Request) (*http.Response, error) {
			resp := &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader(`{"status":"error","code":400,"message":"Invalid parameters: site"}`)),
				Header:     make(http.Header),
				StatusCode: http.StatusBadRequest,
			}
			resp.Header.Add("Content-Type", "application/json; charset=utf-8")
			return resp, nil
		}

		hc := &http.Client{Transport: mockT}

		s, err := New("/", hc, nil)
		if err != nil {
			t.Fatalf("New: unexpected error: %v", err)
		}

		_, err = s.Do(&Query{})

		if err == nil {
			t.Fatal("Search: got nil, want err")
		}

		if _, ok := err.(*APIError); !ok {
			t.Fatalf("error is a %T (%q), want a %T", err, err, &APIError{})
		}

		ae := err.(*APIError)

		if got, want := ae.Error(), "search-api: 400 Invalid parameters: site"; got != want {
			t.Errorf("ae.Error() = %q, want %q", got, want)
		}
	})
}

func TestLogf(t *testing.T) {
	var buf bytes.Buffer
	logf := func(format string, v ...interface{}) {
		fmt.Fprintf(&buf, format, v...)
	}
	s, err := New("/", nil, logf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	s.logf("foo %s", "bar")

	if got, want := buf.String(), "foo bar"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
