package search

import "fmt"

// APIError holds an error as received from the search service.
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("search-api: %d %s", e.Code, e.Message)
}
