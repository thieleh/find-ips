package main

import (
	"find-ip/app"
	"log"
	"os"
)

func main() {

	application := app.Generate()
	error := application.Run(os.Args)

	if error != nil {
		log.Fatal(error)
	}
}
