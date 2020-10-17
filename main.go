package main

import (
	"fmt"
	"log"
	"os"
	"search-ip/app"
)

func main() {
	fmt.Println("🚀 Starting application...")

	application := app.Generate()

	if applicationError := application.Run(os.Args); applicationError != nil {
		log.Fatal(applicationError)
	}

}
