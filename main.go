package main

import (
	"log"
	"os"

	"github.com/leoneville/linha-de-comando/app"
)

func main() {
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
