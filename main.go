package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	r := router.Generate()

	fmt.Printf("Listening on gateway %d", config.Gateway)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Gateway), r))
}
