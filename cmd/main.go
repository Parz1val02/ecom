package main

import (
	"log"

	"github.com/Parz1val02/ecom/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
