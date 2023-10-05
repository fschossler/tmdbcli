package internal

import "os"

var TMDB_CLI_BEARER_TOKEN string = os.Getenv("TMDB_CLI_BEARER_TOKEN")

func ValidateBearerToken() string {
	if TMDB_CLI_BEARER_TOKEN == "" {
		panic("You need to create your API Key and put your your bearer token in the environment variable 'TMDB_CLI_BEARER_TOKEN'. Check more infos on how to do this in the docs.")
	}

	return TMDB_CLI_BEARER_TOKEN
}
