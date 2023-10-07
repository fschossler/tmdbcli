package serie

import (
	"fmt"

	"github.com/fschossler/tmdbcli/cmd"
	"github.com/spf13/cobra"
)

var SerieCmd = &cobra.Command{
	Use:   "serie",
	Short: "Some informations about series",
	Long:  `Some informations about series. You can check for top-rated and top-popular series, synopsis and details about a serie.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serie called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(SerieCmd)
}
