package usecases

import "github.com/fabiohsgomes/go-expert-labs-deploy/internal/infra/clients"

type CalculaTemparaturasUseCase struct {
	weatherapiClient clients.WeatherClient
	consultaCepUseCase *ConsultaCepUseCase
}

func NewCalculaTemperaturasUseCase(weatherapiClient clients.WeatherClient, consultaCepUseCase *ConsultaCepUseCase) *CalculaTemparaturasUseCase {
	return &CalculaTemparaturasUseCase{
		weatherapiClient: weatherapiClient,
		consultaCepUseCase: consultaCepUseCase,
	}
}

type DadosTemperaturas struct {
	Celcius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func (u *CalculaTemparaturasUseCase) Execute(cep string) (*DadosTemperaturas, error) {
	dadosCep, err := u.consultaCepUseCase.ConsultaCep(cep)
	if err != nil {
		return nil, err
	}
	
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
