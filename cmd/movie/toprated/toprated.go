package toprated

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fschossler/tmdbcli/cmd/movie"
	"github.com/fschossler/tmdbcli/internal"
	"github.com/spf13/cobra"
)

var topRatedCmd = &cobra.Command{
	Use:   "toprated",
	Short: "Shows Top Rated movies in TMDB database",
	Long:  `Shows Top Rated movies in TMDB database.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := TopRated()
		if err != nil {
			log.Fatal(err)
		}
	},
}

type Root struct {
	Page    int
	Results []Results
}

type Results struct {
	OriginalTitle string  `json:"original_title"`
	VoteAverage   float32 `json:"vote_average"`
	Overview      string  `json:"overview"`
}

func TopRated() error {

	BEARER_TOKEN := internal.ValidateBearerToken()

	url := "https://api.themoviedb.org/3/movie/top_rated?language=en-US&page=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+BEARER_TOKEN)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var movie Root

	jsonText := string(body)

	err = json.Unmarshal([]byte(jsonText), &movie)
	if err != nil {
		return err
	}

	for _, value := range movie.Results {
		fmt.Println(value.OriginalTitle+":", value.VoteAverage)
		fmt.Println(value.Overview)
	}

	return nil
}

func init() {
	movie.MovieCmd.AddCommand(topRatedCmd)
}
