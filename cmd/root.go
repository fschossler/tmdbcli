package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
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
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
