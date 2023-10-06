package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	Language string
	Page     int
)

var RootCmd = &cobra.Command{
	Use:   "tmdbcli",
	Short: "A CLI created to get infos about movies and series in TMDB database 📽️",
	Long: `A CLI created to get infos about movies and series in TMDB database 📽️. 
	
With 'tmdbcli' you can get infos about: 
- Top rated movies and series
- Top popular movies and series
- Movie infos
- And more.`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&Language, "language", "l", "en-US", "Output language of the CLI. Use the \"Language Code and Country Code\" pattern. Remember to put between quotes. Example: 'en-US'.")
	RootCmd.PersistentFlags().IntVarP(&Page, "page", "p", 1, "Number of the page to return the results. Every page show 20 results.")
}
