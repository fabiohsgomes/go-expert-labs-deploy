package main

import (
	"log"
	"net/http"

	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/config"
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/handlers"
)

func main() {
	config.LoadConfig(".")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /cidades/{cep}/temperaturas", handlers.ProcessaTemperaturasHandler)

	log.Println("Servidor escutando na porta :3000")

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Println(err)
	}
}
