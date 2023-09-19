package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

const version string = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of the tmdbcli",
	Long:  "Version of the tmdbcli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version of the tmdbli:", version)
	},
}
