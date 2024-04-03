package main

import (
	"fmt"
	"linhadecomando/app"
	"log"
	"os"
)



func main (){

	fmt.Println("Ponto de Partida")

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil{
		log.Fatal(erro)
	}
}
