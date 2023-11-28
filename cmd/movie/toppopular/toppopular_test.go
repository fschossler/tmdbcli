// toppopular_test.go
package toppopular

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fschossler/tmdbcli/internal"
	"github.com/stretchr/testify/assert"
)

func TestTopPopular(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request
		assert.Equal(t, "/movie/popular", r.URL.Path, "Expected request to /movie/popular")
		assert.Equal(t, "application/json", r.Header.Get("Accept"), "Expected Accept: application/json header")
		w.WriteHeader(http.StatusOK)
		// Provide mock JSON response
		w.Write([]byte(`{"page": 1, "results": [{"title": "Movie 1", "vote_average": 8.5, "overview": "Overview of Movie 1"}]}`))
	}))
	defer server.Close()

	// Set the base URL for the internal package to the mock server's URL
	internal.SetBaseURL(server.URL)

	// Test the TopPopular function
	err := TopPopular()
	assert.NoError(t, err)
}
