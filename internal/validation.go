package internal

import (
	"os"

	"github.com/joho/godotenv"
)

func ValidateBearerToken() string {

	err := godotenv.Load()
	if err != nil {
		return "Error loading .env file."
	}

	tmdbToken := os.Getenv("TMDB_CLI_BEARER_TOKEN")

	if tmdbToken == "" {
		panic("You need to create your API Key and put your your bearer token in the environment variable 'TMDB_CLI_BEARER_TOKEN'. Check more infos on how to do this in the docs.")
	}

	return tmdbToken
}
