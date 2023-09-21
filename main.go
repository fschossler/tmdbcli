package main

import (
	"github.com/fschossler/tmdbcli/cmd"
	_ "github.com/fschossler/tmdbcli/cmd/movie"
	_ "github.com/fschossler/tmdbcli/cmd/movie/toprated"
	_ "github.com/fschossler/tmdbcli/cmd/serie"
	_ "github.com/fschossler/tmdbcli/cmd/serie/toprated"
	_ "github.com/fschossler/tmdbcli/cmd/version"
)

func main() {
	cmd.Execute()
}
