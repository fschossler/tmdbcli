package toprated

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fschossler/tmdbcli/internal"
	"github.com/stretchr/testify/assert"
)

func TestTopRated(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/tv/top_rated", r.URL.Path, "Expected request to /tv/top_rated")
		assert.Equal(t, "application/json", r.Header.Get("Accept"), "Expected Accept: application/json header")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"page": 1, "results": [{"name": "Serie 1", "vote_average": 8.5, "overview": "Overview of Serie 1"}]}`))
	}))
	defer server.Close()

	// Set the base URL for the internal package to the mock server's URL
	internal.SetBaseURL(server.URL)

	err := TopRated()
	assert.NoError(t, err)
}
