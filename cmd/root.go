package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Language string

var RootCmd = &cobra.Command{
	Use:   "tmdbcli",
	Short: "A CLI created to get infos about movies in TMDB database.",
	Long: `A CLI created to get infos about movies in TMDB database. 
	
With 'tmdbcli' you can get infos about: 
- Top rated movies
- Top popular movies
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
}
