package toprated

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/fschossler/tmdbcli/cmd/serie"
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

var BEARER_TOKEN string = os.Getenv("BEARER_TOKEN")

type Root struct {
	Page    int
	Results []Results
}

type Results struct {
	OriginalName string  `json:"original_name"`
	VoteAverage  float32 `json:"vote_average"`
}

func TopRated() error {

	if BEARER_TOKEN == "" {
		panic("You need to create your API Key and put your your bearer token in the environment variable 'BEARER_TOKEN'. Check more infos on how to do this in the docs.")
	}

	url := "https://api.themoviedb.org/3/tv/top_rated?language=en-US&page=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+BEARER_TOKEN)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var serie Root

	jsonText := string(body)

	err = json.Unmarshal([]byte(jsonText), &serie)
	if err != nil {
		return err
	}

	for _, value := range serie.Results {
		fmt.Println(value.OriginalName+":", value.VoteAverage)
	}

	return nil
}

func init() {
	serie.SerieCmd.AddCommand(topRatedCmd)
}
