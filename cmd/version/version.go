package version

import (
	"fmt"

	"github.com/fschossler/tmdbcli/cmd"
	"github.com/spf13/cobra"
)

const version string = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of the tmdbcli",
	Long:  "Version of the tmdbcli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version of the tmdbcli:", version)
	},
}

func init() {
	cmd.RootCmd.AddCommand(versionCmd)
}
