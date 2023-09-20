package movie

import (
	"fmt"

	"github.com/fschossler/tmdbcli/cmd"
	"github.com/spf13/cobra"
)

var MovieCmd = &cobra.Command{
	Use:   "movie",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("movie called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(MovieCmd)
}
