package usecases

import (
	"testing"

	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/domain"
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/erros"
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/infra/clients"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ConsultaCepTestSuite struct {
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

func TestConsultaCepSuite(t *testing.T) {
	suite.Run(t, new(ConsultaCepTestSuite))
}

func (s *ConsultaCepTestSuite) TestConsultaCepWithValidCep() {
	// Arrange
	cepClientMock := new(ViaCepClientMock)
	consultaCepUseCase := NewConsultaCepUseCase(cepClientMock)

	cep, _ := domain.NewCep("01001000")
	expectedResponse := &clients.DadosCepResponse{
		Cep:         "01001000",
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Bairro:      "Sé",
		Localidade:  "São Paulo",
		Uf:          "SP",
	}

	cepClientMock.On("ConsultaCep", cep.Codigo()).Return(expectedResponse, nil)

	// Act
	dadosCep, err := consultaCepUseCase.ConsultaCep(cep)

	// Assert
	s.NoError(err)
	s.Equal(expectedResponse.Cep, dadosCep.Cep)
	s.Equal(expectedResponse.Logradouro, dadosCep.Logradouro)
	s.Equal(expectedResponse.Complemento, dadosCep.Complemento)
	s.Equal(expectedResponse.Bairro, dadosCep.Bairro)
	s.Equal(expectedResponse.Localidade, dadosCep.Localidade)
	s.Equal(expectedResponse.Uf, dadosCep.Uf)

	cepClientMock.AssertExpectations(s.T())
}

func (s *ConsultaCepTestSuite) TestConsultaCepWithUnexistCep() {
	// Arrange
	cepClientMock := new(ViaCepClientMock)
	consultaCepUseCase := NewConsultaCepUseCase(cepClientMock)

	cep, _ := domain.NewCep("00000000")
	expectedResponse := &clients.DadosCepResponse{
		Erro: "true",
	}
	expectedResponseError := erros.ErrZipCodeNotFound

	cepClientMock.On("ConsultaCep", cep.Codigo()).Return(expectedResponse, expectedResponseError)

	// Act
	_, err := consultaCepUseCase.ConsultaCep(cep)

	// Assert
	s.Error(err)
	s.ErrorIs(err, erros.ErrZipCodeNotFound)
	s.Equal(expectedResponseError.Error(), err.Error())

	cepClientMock.AssertExpectations(s.T())
}
