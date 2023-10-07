package movie

import (
	"fmt"

	"github.com/fschossler/tmdbcli/cmd"
	"github.com/spf13/cobra"
)

var MovieCmd = &cobra.Command{
	Use:   "movie",
	Short: "Some informations about movies",
	Long:  `Some informations about movies. You can check for Top Rated and Top Popular movies, synopsis and details about a movie.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Check the possible subcommands with the \"tmdbcli movie --help\".")
	},
}

func init() {
	cmd.RootCmd.AddCommand(MovieCmd)
}
