package app

import (
	"github.com/urfave/cli"
)

// Generate returns the command line application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Search IP"
	app.Usage = "Search IP's and web servers"

	return app
}
