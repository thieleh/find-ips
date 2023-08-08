package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return application ready to be executed
func Generate() *cli.App {

	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search IPs and Server names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "neopets.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search Ips",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "servers",
			Usage:  "Find the name of the servers",
			Flags:  flags,
			Action: findServers,
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

func findServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

/*
read-me to be created...
to execute the application and find the ip:
go run main.go ip --host [flag-host]

to execute the application and find the server:
go run main.go servers --host [flag-host]
*/
