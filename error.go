package search

import (
	"fmt"
	"net/http"
)

// APIError holds an error as received from the search service.
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("search-api: HTTP %d %s: %s", e.Code, http.StatusText(e.Code), e.Message)
}
