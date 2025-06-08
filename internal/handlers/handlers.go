package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/erros"
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/infra/clients"
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/service"
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/usecases"
)

func ProcessaTemperaturasHandler(w http.ResponseWriter, r *http.Request) {
	cepPathValue := r.PathValue("cep")
	viaCepClient := clients.NewViaCepClient()
	cepUseCase := usecases.NewConsultaCepUseCase(viaCepClient)

	weatherApiClient := clients.NewWeatherApiClient()
	calculaTemperaturasUseCase := usecases.NewCalculaTemperaturasUseCase(weatherApiClient)

	temperaturasService := service.NewTemperaturasService(cepUseCase, calculaTemperaturasUseCase)
	dadosTemperaturas, err := temperaturasService.Processa(cepPathValue)
	if err != nil {
		if errors.Is(err, erros.ErrInvalidZipCode) || errors.Is(err, erros.ErrCityIsRequired) {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if errors.Is(err, erros.ErrZipCodeNotFound) || errors.Is(err, erros.ErrCityNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error processing temperatures for CEP %s: %v", cepPathValue, err)
		return
	}

	http.Header.Set(w.Header(), "Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dadosTemperaturas); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error encoding response for CEP %s: %v", cepPathValue, err)
		return
	}
	log.Printf("Successfully processed temperatures for CEP %s", cepPathValue)
}
