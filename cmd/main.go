package main

import (
	"api/internal/config"
	"api/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Load()
	r := router.Generate()

	fmt.Printf("Listen in the port %d", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
