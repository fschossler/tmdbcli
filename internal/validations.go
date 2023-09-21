package internal

import "os"

var BEARER_TOKEN string = os.Getenv("BEARER_TOKEN")

func ValidateBearerToken() string {
	if BEARER_TOKEN == "" {
		panic("You need to create your API Key and put your your bearer token in the environment variable 'BEARER_TOKEN'. Check more infos on how to do this in the docs.")
	}

	return BEARER_TOKEN
}
