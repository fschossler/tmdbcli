package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var topRatedCmd = &cobra.Command{
	Use:   "top-rated",
	Short: "Shows Top Rated movies in TMDB database",
	Long:  `Shows Top Rated movies in TMDB database.`,
	Run: func(cmd *cobra.Command, args []string) {
		TopRated()
	},
}

func init() {
	rootCmd.AddCommand(topRatedCmd)
}

var BEARER_TOKEN string = os.Getenv("BEARER_TOKEN")

type Root struct {
	Page    int
	Results []Results
}

type Results struct {
	OriginalTitle string  `json:"original_title"`
	VoteAverage   float32 `json:"vote_average"`
}

func TopRated() {

	if BEARER_TOKEN == "" {
		// fmt.Println("You need to create your API Key and put your your bearer token in the environment variable 'BEARER_TOKEN'.")
		panic("You need to create your API Key and put your your bearer token in the environment variable 'BEARER_TOKEN'. Check more infos on how to do this in the docs.")
	}

	url := "https://api.themoviedb.org/3/movie/top_rated?language=en-US&page=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+BEARER_TOKEN)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var movies Root

	jsonText := string(body)

	err2 := json.Unmarshal([]byte(jsonText), &movies)
	if err2 != nil {
		fmt.Println(err2)
	}

	for _, value := range movies.Results {
		fmt.Println(value.OriginalTitle+":", value.VoteAverage)
	}

}
