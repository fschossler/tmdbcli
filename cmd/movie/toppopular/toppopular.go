package toppopular

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/fschossler/tmdbcli/cmd/movie"
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
	Title       string  `json:"title"`
	VoteAverage float32 `json:"vote_average"`
	Overview    string  `json:"overview"`
}

func TopPopular() error {

	jsonReturn := internal.RequestPath("/movie/popular")
	var movie Root

	err := json.Unmarshal([]byte(jsonReturn), &movie)
	if err != nil {
		return err
	}

	for _, movie := range movie.Results {
		title := color.New(color.FgHiCyan)
		voteAverage := color.New(color.FgGreen)

		title.Print(movie.Title + ": ")
		voteAverage.Println(movie.VoteAverage)
		fmt.Println(movie.Overview)
		fmt.Println("")
	}

	return nil
}

func init() {
	movie.MovieCmd.AddCommand(topPopularCmd)
}
