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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br", //default value
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IP's when recieve a url",
			Flags:  flags,
			Action: searchIPByHost,
		},
		{
			Name:   "servers",
			Usage:  "Search the server name",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIPByHost(c *cli.Context) {
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

func searchServers(c *cli.Context) {
	fmt.Println("Searching servers...")

	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
