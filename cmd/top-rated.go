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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		TopRated()
	},
}

func init() {
	rootCmd.AddCommand(topRatedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topRatedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topRatedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
