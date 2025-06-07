package service

import (
	"fmt"
	"testing"

	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/erros"
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/helpers"
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/infra/clients"
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/usecases"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TemperaturasServiceTestSuite struct {
	suite.Suite
}

type ViaCepClientMock struct {
	mock.Mock
}

func (m *ViaCepClientMock) ConsultaCep(cep string) (*clients.DadosCepResponse, error) {
	args := m.Called(cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*clients.DadosCepResponse), args.Error(1)
}

type WeatherApiClientMock struct {
	mock.Mock
}

func (m *WeatherApiClientMock) ConsultaClima(cidade string) (*clients.WeatherResponse, error) {
	args := m.Called(cidade)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*clients.WeatherResponse), args.Error(1)
}

func TestTemperaturasServiceSuite(t *testing.T) {
	suite.Run(t, new(TemperaturasServiceTestSuite))
}

func (s *TemperaturasServiceTestSuite) ProcessaTemperaturasComCepValido() {
	// Arrange
	viacepClientMock := new(ViaCepClientMock)
	consultaCepUseCase := usecases.NewConsultaCepUseCase(viacepClientMock)

	weatherapiClientMock := new(WeatherApiClientMock)
	calculaTemperaturasUseCase := usecases.NewCalculaTemperaturasUseCase(weatherapiClientMock)

	service := NewTemperaturasService(consultaCepUseCase, calculaTemperaturasUseCase)

	//expected DadosCepResponse
	dadosCepResponseMock := &clients.DadosCepResponse{
		Cep:         "01001000",
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Bairro:      "Sé",
		Localidade:  "São Paulo",
		Uf:          "SP",
	}

	//expected WeatherResponse with only TempC
	weatherResponseMock := &clients.WeatherResponse{
		Current: clients.Current{
			TempC: 25.0,
		},
	}

	expectedCelsius := fmt.Sprintf("%.1f", weatherResponseMock.Current.TempC)
	expectedFahrenheit := fmt.Sprintf("%.1f", helpers.CelsiusToFahrenheit(weatherResponseMock.Current.TempC))
	expectedKelvin := fmt.Sprintf("%.1f", helpers.CelsiusToKelvin(weatherResponseMock.Current.TempC))

	// Mocking the expected behavior
	viacepClientMock.On("ConsultaCep", "01001000").Return(dadosCepResponseMock, nil)
	weatherapiClientMock.On("ConsultaClima", "São Paulo").Return(weatherResponseMock, nil)

	// Act
	dadosTemperaturas, err := service.Processa("01001000")

	// Assert
	s.NoError(err)
	s.Equal(expectedCelsius, dadosTemperaturas.Celcius)
	s.Equal(expectedFahrenheit, dadosTemperaturas.Fahrenheit)
	s.Equal(expectedKelvin, dadosTemperaturas.Kelvin)

	viacepClientMock.AssertExpectations(s.T())
	weatherapiClientMock.AssertExpectations(s.T())
}

func (s *TemperaturasServiceTestSuite) ProcessaTemperaturasComCepInexistente() {
	// Arrange
	viacepClientMock := new(ViaCepClientMock)
	consultaCepUseCase := usecases.NewConsultaCepUseCase(viacepClientMock)

	weatherapiClientMock := new(WeatherApiClientMock)
	calculaTemperaturasUseCase := usecases.NewCalculaTemperaturasUseCase(weatherapiClientMock)

	service := NewTemperaturasService(consultaCepUseCase, calculaTemperaturasUseCase)

	//expected DadosCepResponse
	dadosCepResponseMock := &clients.DadosCepResponse{
		Erro: "true",
	}

	expectedErrDadosCepResponse := erros.ErrZipCodeNotFound

	// Mocking the expected behavior
	viacepClientMock.On("ConsultaCep", "00000000").Return(dadosCepResponseMock, expectedErrDadosCepResponse)

	// Act
	_, err := service.Processa("00000000")

	// Assert
	s.Error(err)
	s.ErrorIs(err, erros.ErrZipCodeNotFound)
	s.Equal(expectedErrDadosCepResponse.Error(), err.Error())

	viacepClientMock.AssertExpectations(s.T())
	weatherapiClientMock.AssertNotCalled(s.T(), "ConsultaClima", mock.Anything)
}

func (s *TemperaturasServiceTestSuite) ProcessaTemperaturasComCepInvalido() {
	// Arrange
	viacepClientMock := new(ViaCepClientMock)
	consultaCepUseCase := usecases.NewConsultaCepUseCase(viacepClientMock)

	weatherapiClientMock := new(WeatherApiClientMock)
	calculaTemperaturasUseCase := usecases.NewCalculaTemperaturasUseCase(weatherapiClientMock)

	service := NewTemperaturasService(consultaCepUseCase, calculaTemperaturasUseCase)

	expectedErr := erros.ErrZipCodeNotFound

	// Act
	_, err := service.Processa("08931a30")

	// Assert
	s.Error(err)
	s.ErrorIs(err, erros.ErrInvalidZipCode)
	s.Equal(expectedErr.Error(), err.Error())

	viacepClientMock.AssertNotCalled(s.T(), "ConsultaCep", mock.Anything)
	weatherapiClientMock.AssertNotCalled(s.T(), "ConsultaClima", mock.Anything)
}