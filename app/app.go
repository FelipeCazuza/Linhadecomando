package app

import (
	
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"

)

//Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Buscas IPS e Nomes de Servidor na internet"

	flags :=  []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "goolge.com.br",
		},
		
	}

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Buscar IPS de Endereço na Internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name:  "servidores",
			Usage: "buscar o nome do servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app

}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context){
	host := c.String("host")

	Servidores, erro := net.LookupNS(host) //Nameserver
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range Servidores{
		fmt.Println(servidor.Host)
	}
}
