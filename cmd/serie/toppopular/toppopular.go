package toppopular

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/fschossler/tmdbcli/cmd/serie"
	"github.com/fschossler/tmdbcli/internal"
	"github.com/spf13/cobra"
)

var topPopularCmd = &cobra.Command{
	Use:   "toppopular",
	Short: "Shows Top Popular movies in TMDB database",
	Long:  `Shows Top Popular movies in TMDB database.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := TopPopular()
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
	Name        string  `json:"name"`
	VoteAverage float32 `json:"vote_average"`
	Overview    string  `json:"overview"`
}

func TopPopular() error {

	BEARER_TOKEN := internal.ValidateBearerToken()

	url := "https://api.themoviedb.org/3/tv/popular?language=en-US&page=1"

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

	for _, movie := range movie.Results {
		title := color.New(color.FgHiCyan)
		voteAverage := color.New(color.FgGreen)

		title.Print(movie.Name + ": ")
		voteAverage.Println(movie.VoteAverage)
		fmt.Println(movie.Overview)
		fmt.Println("")
	}

	return nil
}

func init() {
	serie.SerieCmd.AddCommand(topPopularCmd)
}