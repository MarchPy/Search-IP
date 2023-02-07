package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Search IP"
	app.Usage = "Busca de IPs e nomes de servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de IPs e nomes de servidores na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca de nome de servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app

}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}
	for _, ips := range ips {
		fmt.Println(ips)
	}

}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, servidores := range servidores {
		fmt.Println(servidores.Host)
	}
}
