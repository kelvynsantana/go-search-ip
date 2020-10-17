package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns the command line application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Search IP"
	app.Usage = "Search IP's and web servers"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IP's when recieve a url",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com.br", //default value
				},
			},
			Action: searchByIP,
		},
	}

	return app
}

func searchByIP(c *cli.Context) {
	fmt.Println("Searching IP's...")

	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
