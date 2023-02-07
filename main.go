package main

import (
	"log"
	"os"
	app "search-ip/App"
)

func main() {
	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
