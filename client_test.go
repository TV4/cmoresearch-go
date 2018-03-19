package search

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("DefaultConfig", func(t *testing.T) {
		c := NewClient()

		if got, want := c.baseURL.String(), "https://cmore-search.b17g.services"; got != want {
			t.Errorf("c.baseURL.String() = %q, want %q", got, want)
		}

		if c.debugLogf == nil {
			t.Errorf("c.debugLogf is nil")
		}
	})

	t.Run("SetBaseURL", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			c := NewClient(SetBaseURL("http://example.com/"))

			if got, want := c.baseURL.String(), "http://example.com/"; got != want {
				t.Errorf("c.baseURL.String() = %q, want %q", got, want)
			}
		})

		t.Run("Fail", func(t *testing.T) {
			c := NewClient(SetBaseURL(": not an URL"))

			if c.baseURL != nil {
				t.Error("c.baseURL is not nil")
			}
		})
	})

	t.Run("SetHTTPClient", func(t *testing.T) {
		hc := &http.Client{}

		c := NewClient(SetBaseURL("/"), SetHTTPClient(hc))

		if got, want := c.httpClient, hc; got != want {
			t.Errorf("c.httpClient = %p, want %p", got, want)
		}
	})

	t.Run("SetDebugLogf", func(t *testing.T) {
		logged := ""
		logf := func(format string, v ...interface{}) {
			logged = fmt.Sprintf(format, v...)
		}
		c := NewClient(SetDebugLogf(logf))
		c.debugLogf("foo %s", "bar")

		if got, want := logged, "foo bar"; got != want {
			t.Errorf("c.debugLogf not properly configured (logged %q, want %q)", got, want)
		}
	})
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
