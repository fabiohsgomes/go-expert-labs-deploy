package main

import (
	"net/http"

	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cidades/{cep}/temperaturas", handlers.ProcessaTemperaturasHandler)

	http.ListenAndServe(":3000", mux)
}
