package usecases

import "github.com/fabiohsgomes/go-expert-labs-deploy/internal/infra/clients"

type CalculaTemparaturasUseCase struct {
	weatherapiClient clients.WeatherClient
}

func NewCalculaTemperaturasUseCase(weatherapiClient clients.WeatherClient) *CalculaTemparaturasUseCase {
	return &CalculaTemparaturasUseCase{
		weatherapiClient: weatherapiClient,
	}
}

type DadosTemperaturas struct {
	Celcius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func (u *CalculaTemparaturasUseCase) Execute(dadosCep *DadosCep) (*DadosTemperaturas, error) {
	weatherResponse, err := u.weatherapiClient.ConsultaClima(dadosCep.Localidade)
	if err != nil {
		return nil, err
	}

	dadosTemperaturas := &DadosTemperaturas{
		Celcius:    weatherResponse.Current.TempC,
		Fahrenheit: weatherResponse.Current.TempF,
		Kelvin:     0.0,
	}

	return dadosTemperaturas, nil
}
