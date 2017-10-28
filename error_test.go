package search

import (
	"net/http"
	"testing"
)

func TestAPIError_Error(t *testing.T) {
	err := &APIError{
		Code:    http.StatusTeapot,
		Message: "Foo message",
	}

	if got, want := err.Error(), "search-api: HTTP 418 I'm a teapot: Foo message"; got != want {
		t.Errorf("err.Error() = %q, want %q", got, want)
	}
}
