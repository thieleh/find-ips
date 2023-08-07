package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return aplciation ready to be executed
func Generate() *cli.App {

	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search IPs and Server names"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search Ips",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com/thieleh",
				},
			},
			Action: findIps,
		},
	}

	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
