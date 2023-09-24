package main

import (
	"log"
	"net/http"

	"github.com/alcb1310/go-api/internals/api"
	"github.com/gorilla/handlers"
)

func main() {
	srv := api.NewServer()

	originsOK := handlers.AllowedOrigins([]string{
		"*",
	})
	methodsOK := handlers.AllowedHeaders([]string{
		http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodDelete,
		http.MethodPatch,
	})

	log.Panic(http.ListenAndServe(":42069", handlers.CORS(originsOK, methodsOK)(srv)))
}
