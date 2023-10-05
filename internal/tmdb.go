package internal

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/fschossler/tmdbcli/cmd"
)

const TmdbBaseUrl = "https://api.themoviedb.org/3"

func RequestPath(path string) string {

	TMDB_CLI_BEARER_TOKEN := ValidateBearerToken()

	baseUrl := TmdbBaseUrl
	path += "?"
	fullUrl := "" + baseUrl + "" + path + ""

	req, _ := http.NewRequest("GET", fullUrl, nil)

	// Headers
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+TMDB_CLI_BEARER_TOKEN)

	// Query Parameters to change the request behavior
	queryParameters := req.URL.Query()
	queryParameters.Add("language", cmd.Language)
	queryParameters.Add("page", strconv.Itoa(cmd.Page))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error in the HTTP request.")
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	jsonText := string(body)

	return jsonText
}
