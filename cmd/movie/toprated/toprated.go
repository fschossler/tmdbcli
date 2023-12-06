package toprated

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fatih/color"
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
	Title       string  `json:"title"`
	VoteAverage float32 `json:"vote_average"`
	Overview    string  `json:"overview"`
}

func TopRated() error {

	jsonReturn := internal.RequestPath(internal.UrlParams{Path: "/movie/top_rated"})

	var movie Root

	err := json.Unmarshal([]byte(jsonReturn), &movie)
	if err != nil {
		return err
	}

	for _, movie := range movie.Results {
		title := color.New(color.FgHiCyan)
		voteAverage := color.New(color.FgGreen)

		title.Print(movie.Title + ": ")
		voteAverage.Printf("%.1f\n", movie.VoteAverage)
		fmt.Println(movie.Overview)
		fmt.Println("")
	}

	return nil
}

func init() {
	movie.MovieCmd.AddCommand(topRatedCmd)
}
