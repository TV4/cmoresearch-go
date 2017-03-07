package search

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("DefaultConfig", func(t *testing.T) {
		c := NewClient()

		if got, want := c.baseURL.String(), "https://search.b17g.services"; got != want {
			t.Errorf("s.baseURL.String() = %q, want %q", got, want)
		}
	})

	t.Run("SetBaseURL", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			c := NewClient(SetBaseURL("http://example.com/"))

			if got, want := c.baseURL.String(), "http://example.com/"; got != want {
				t.Errorf("s.baseURL.String() = %q, want %q", got, want)
			}
		})

		t.Run("Fail", func(t *testing.T) {
			c := NewClient(SetBaseURL(": not an URL"))

			if c.baseURL != nil {
				t.Error("s.baseURL is not nil")
			}
		})
	})

	t.Run("SetLogf", func(t *testing.T) {
		var buf bytes.Buffer
		logf := func(format string, v ...interface{}) {
			fmt.Fprintf(&buf, format, v...)
		}
		c := NewClient(SetBaseURL("/"), SetLogf(logf))

		c.logf("foo %s", "bar")

		if got, want := buf.String(), "foo bar"; got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("SetHTTPClient", func(t *testing.T) {
		hc := &http.Client{}

		c := NewClient(SetBaseURL("/"), SetHTTPClient(hc))

		if got, want := c.httpClient, hc; got != want {
			t.Errorf("s.httpClient = %p, want %p", got, want)
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
