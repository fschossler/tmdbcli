package toprated

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/fschossler/tmdbcli/cmd/serie"
	"github.com/fschossler/tmdbcli/internal"
	"github.com/spf13/cobra"
)

var topRatedCmd = &cobra.Command{
	Use:   "toprated",
	Short: "Shows Top Rated series in TMDB database",
	Long:  `Shows Top Rated series in TMDB database.`,
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
	Name        string  `json:"name"`
	VoteAverage float32 `json:"vote_average"`
	Overview    string  `json:"overview"`
}

func TopRated() error {

	jsonReturn := internal.RequestPath("/tv/top_rated")

	var serie Root

	err := json.Unmarshal([]byte(jsonReturn), &serie)
	if err != nil {
		return err
	}

	for _, serie := range serie.Results {
		title := color.New(color.FgHiCyan)
		voteAverage := color.New(color.FgGreen)

		title.Print(serie.Name + ": ")
		voteAverage.Printf("%.1f\n", serie.VoteAverage)
		fmt.Println(serie.Overview)
		fmt.Println("")
	}

	return nil
}

func init() {
	serie.SerieCmd.AddCommand(topRatedCmd)
}
