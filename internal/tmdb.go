package internal

import (
	"context"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fschossler/tmdbcli/cmd"
)

type UrlParams struct {
	Path    string
	BaseUrl string
}

var TmdbBaseUrl = "https://api.themoviedb.org/3"

func RequestPath(url UrlParams) string {

	TMDB_CLI_BEARER_TOKEN := ValidateBearerToken()

	if url.BaseUrl != "" {
		TmdbBaseUrl += url.BaseUrl
	} else {
		TmdbBaseUrl += url.Path
	}

	req, _ := http.NewRequest("GET", TmdbBaseUrl, nil)

	// Headers
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+TMDB_CLI_BEARER_TOKEN)

	// Query Parameters to change the request behavior
	queryParameters := req.URL.Query()
	queryParameters.Add("language", cmd.Language)
	queryParameters.Add("page", strconv.Itoa(cmd.Page))
	req.URL.RawQuery = queryParameters.Encode()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error in the HTTP request.")
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	jsonText := string(body)

	return jsonText
}
