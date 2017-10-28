package search

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type mockTransport func(*http.Request) (*http.Response, error)

func (mt mockTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return mt(r)
}

func TestSearch(t *testing.T) {
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

			c := NewClient(SetBaseURL("/"), SetHTTPClient(hc))

			_, err := c.Search(context.Background(), nil)

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

		c := NewClient(SetBaseURL("/"), SetHTTPClient(hc))

		_, err := c.Search(context.Background(), nil)

		if err == nil {
			t.Fatal("Search: got nil, want err")
		}

		if _, ok := err.(*APIError); !ok {
			t.Fatalf("error is a %T (%q), want a %T", err, err, &APIError{})
		}

		ae := err.(*APIError)

		if got, want := ae.Error(), "search-api: HTTP 400 Bad Request: Invalid parameters: site"; got != want {
			t.Errorf("ae.Error() = %q, want %q", got, want)
		}
	})

	t.Run("Meta", func(t *testing.T) {
		var mockT mockTransport = func(r *http.Request) (*http.Response, error) {
			resp := &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(`{}`)),
				Header: http.Header{
					"X-Foo": {"foo-value"},
				},
				StatusCode: http.StatusTeapot,
			}
			resp.Header.Add("Content-Type", "application/json; charset=utf-8")
			return resp, nil
		}

		hc := &http.Client{Transport: mockT}

		c := NewClient(SetBaseURL("/"), SetHTTPClient(hc))

		res, err := c.Search(context.Background(), nil)

		if err == nil {
			t.Fatal("Search: got nil, want err")
		}

		if res.Meta.Header == nil {
			t.Fatalf("res.Meta.Header is nil, want not nil")
		}

		if got, want := res.Meta.Header.Get("X-Foo"), "foo-value"; got != want {
			t.Errorf(`res.Meta.Header.Get("X-Foo") = %q, want %q`, got, want)
		}

		if got, want := res.Meta.StatusCode, http.StatusTeapot; got != want {
			t.Errorf("res.Meta.StatusCode = %d, want %d", got, want)
		}
	})
}

func TestNewSearchRequest(t *testing.T) {
	t.Run("ArbitraryOption", func(t *testing.T) {
		c := NewClient()

		option := func(req *http.Request) {
			req.Header.Set("Foo", "foo-header")
		}

		req, err := c.newSearchRequest(context.Background(), nil, option)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := req.Header.Get("Foo"), "foo-header"; got != want {
			t.Errorf("Foo = %q, want %q", got, want)
		}
	})

	t.Run("SetRequestIDOption", func(t *testing.T) {
		c := NewClient()

		req, err := c.newSearchRequest(context.Background(), nil, SetRequestID("request-id"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := req.Header.Get("X-Request-Id"), "request-id"; got != want {
			t.Errorf("X-Request-Id = %q, want %q", got, want)
		}
	})

	t.Run("QueryString", func(t *testing.T) {
		c := NewClient()

		query := url.Values{}
		query.Add("foo", "123")
		query.Add("bar&", "234 567")
		query.Add("baz", "345")

		req, err := c.newSearchRequest(context.Background(), query)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := req.URL.Query().Encode(), "bar%26=234+567&baz=345&foo=123"; got != want {
			t.Errorf("queryString = %q, want %q", got, want)
		}
	})

	t.Run("FieldsQueryParam", func(t *testing.T) {
		for n, tc := range []struct {
			inputQuery      url.Values
			wantQueryString string
		}{
			{url.Values{}, ""},                                                               // no change
			{url.Values{"fields": {"foo"}}, "fields=foo%2Ctype"},                             // type added
			{url.Values{"fields": {"foo,qux"}}, "fields=foo%2Cqux%2Ctype"},                   // type added
			{url.Values{"fields": {"foo"}, "gurka": {"bar"}}, "fields=foo%2Ctype&gurka=bar"}, // type added
			{url.Values{"fields": {"foo,footype"}}, "fields=foo%2Cfootype%2Ctype"},           // type added
			{url.Values{"fields": {"foo,bar,type"}}, "fields=foo%2Cbar%2Ctype"},              // no change
			{url.Values{"fields": {"foo,type,bar"}}, "fields=foo%2Ctype%2Cbar"},              // no change
			{url.Values{"fields": {"type,foo,bar"}}, "fields=type%2Cfoo%2Cbar"},              // no change
			{url.Values{"fields": {"type"}}, "fields=type"},                                  // no change
		} {
			c := NewClient()

			request, err := c.newSearchRequest(context.Background(), tc.inputQuery)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got, want := request.URL.Query().Encode(), tc.wantQueryString; got != want {
				t.Errorf("[%d] got %q, want %q", n, got, want)
			}
		}
	})
}

func TestMakeResponse(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		target := "http://example.com/foo"

		req := httptest.NewRequest(http.MethodGet, target, nil)

		resp := &http.Response{
			Body: ioutil.NopCloser(strings.NewReader(`
		{
			"total_hits": 100,
			"assets": [
				{ "type": "movie" },
				{ "type": "series" },
				{ "type": "unknown" }
			]
		}
		`)),
			Header:     http.Header{"Foo": {"Bar"}},
			StatusCode: http.StatusTeapot,
		}

		response, err := makeResponse(req, resp)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got, want := response.TotalHits, 100; got != want {
			t.Errorf("response.TotalHits = %d, want %d", got, want)
		}

		if got, want := len(response.Hits), 3; got != want {
			t.Fatalf("len(response.Hits) = %d, want %d", got, want)
		}

		if _, ok := response.Hits[0].(*Asset); !ok {
			t.Errorf("response.Hits[0] is a %T, want a %T", response.Hits[0], &Asset{})
		}

		if _, ok := response.Hits[1].(*Series); !ok {
			t.Errorf("response.Hits[1] is a %T, want a %T", response.Hits[1], &Series{})
		}

		if _, ok := response.Hits[2].(*Asset); !ok {
			t.Errorf("response.Hits[2] is a %T, want a %T", response.Hits[2], &Asset{})
		}

		if got, want := response.Meta.RequestURL.String(), target; got != want {
			t.Errorf(`response.Meta.RequestURL.String() = %q, want %q`, got, want)
		}

		if got, want := response.Meta.Header.Get("Foo"), "Bar"; got != want {
			t.Errorf(`response.Meta.Header.Get("Foo") = %q, want %q`, got, want)
		}

		if got, want := response.Meta.StatusCode, http.StatusTeapot; got != want {
			t.Errorf("response.Meta.StatusCode = %d, want %d", got, want)
		}
	})

	t.Run("Malformed", func(t *testing.T) {
		for _, tc := range []struct {
			description string
			body        string
		}{
			{
				"Type",
				`
					{
						"total_hits": 100,
						"assets": [
							{
								"type": {
									"malformed": "type should not be an object"
								}
							}
						]
					}
				`,
			},
			{
				"Asset",
				`
					{
						"total_hits": 100,
						"assets": [
							{
								"type": "movie",
								"title_sv": {
									"malformed": "title_sv should not be an object"
								}
							}
						]
					}
				`,
			},
			{
				"Series",
				`
					{
						"total_hits": 100,
						"assets": [
							{
								"type": "series",
								"title_sv": {
									"malformed": "title_sv should not be an object"
								}
							}
						]
					}
				`,
			},
		} {
			t.Run(tc.description, func(t *testing.T) {
				target := "http://example.com/" + tc.description

				req := httptest.NewRequest(http.MethodGet, target, nil)

				resp := &http.Response{
					Body:       ioutil.NopCloser(strings.NewReader(tc.body)),
					Header:     http.Header{"Foo": {"Bar"}},
					StatusCode: http.StatusTeapot,
				}

				response, err := makeResponse(req, resp)
				if err == nil {
					t.Fatal("got nil, want error")
				}

				if got, want := response.TotalHits, 100; got != want {
					t.Errorf("response.TotalHits = %d, want %d", got, want)
				}

				if got, want := len(response.Hits), 0; got != want {
					t.Errorf("len(response.Hits) = %d, want %d", got, want)
				}

				if got, want := response.Meta.RequestURL.String(), target; got != want {
					t.Errorf(`response.Meta.RequestURL.String() = %q, want %q`, got, want)
				}

				if got, want := response.Meta.Header.Get("Foo"), "Bar"; got != want {
					t.Errorf(`response.Meta.Header.Get("Foo") = %q, want %q`, got, want)
				}

				if got, want := response.Meta.StatusCode, http.StatusTeapot; got != want {
					t.Errorf("response.Meta.StatusCode = %d, want %d", got, want)
				}
			})
		}
	})

	t.Run("TypeMissing", func(t *testing.T) {
		target := "http://example.com/bar"

		req := httptest.NewRequest(http.MethodGet, target, nil)

		resp := &http.Response{
			Body: ioutil.NopCloser(strings.NewReader(`
				{
					"total_hits": 100,
					"assets": [{}]
				}
		`)),
			Header:     http.Header{"Foo": {"Bar"}},
			StatusCode: http.StatusTeapot,
		}

		response, err := makeResponse(req, resp)
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if got, want := response.TotalHits, 100; got != want {
			t.Errorf("response.TotalHits = %d, want %d", got, want)
		}

		if got, want := len(response.Hits), 0; got != want {
			t.Errorf("len(response.Hits) = %d, want %d", got, want)
		}

		if got, want := response.Meta.RequestURL.String(), target; got != want {
			t.Errorf(`response.Meta.RequestURL.String() = %q, want %q`, got, want)
		}

		if got, want := response.Meta.Header.Get("Foo"), "Bar"; got != want {
			t.Errorf(`response.Meta.Header.Get("Foo") = %q, want %q`, got, want)
		}

		if got, want := response.Meta.StatusCode, http.StatusTeapot; got != want {
			t.Errorf("response.Meta.StatusCode = %d, want %d", got, want)
		}
	})

	t.Run("BodyNotJSON", func(t *testing.T) {
		target := "http://example.com/baz"

		req := httptest.NewRequest(http.MethodGet, target, nil)

		resp := &http.Response{
			Body:       ioutil.NopCloser(strings.NewReader("not-json")),
			Header:     http.Header{"Foo": {"Bar"}},
			StatusCode: http.StatusTeapot,
		}

		response, err := makeResponse(req, resp)
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if got, want := response.Meta.RequestURL.String(), target; got != want {
			t.Errorf(`response.Meta.RequestURL.String() = %q, want %q`, got, want)
		}

		if got, want := response.Meta.Header.Get("Foo"), "Bar"; got != want {
			t.Errorf(`response.Meta.Header.Get("Foo") = %q, want %q`, got, want)
		}

		if got, want := response.Meta.StatusCode, http.StatusTeapot; got != want {
			t.Errorf("response.Meta.StatusCode = %d, want %d", got, want)
		}
	})
}
