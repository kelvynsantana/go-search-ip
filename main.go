package main

import (
	"fmt"
	"log"
	"os"
	"search-ip/app"
)

func main() {
	fmt.Println("Iniciando a aplicação")

	application := app.Generate()

	if applicationError := application.Run(os.Args); applicationError != nil {
		log.Fatal(applicationError)
	}

}
