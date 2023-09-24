package main

import (
	"log"
	"net/http"

	"github.com/alcb1310/go-api/api"
)

func main() {
	srv := api.NewServer()

	log.Panic(http.ListenAndServe(":42069", srv))
}
