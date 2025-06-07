package clients

import (
	"testing"

	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/erros"
	"github.com/stretchr/testify/assert"
)

func TestConsultaCep(t *testing.T) {
	t.Run("Consulta Cep valido", func(t *testing.T) {
		//Arrange
		client := NewViaCepClient()
		cep := "05893130"

		//Act
		dadosCep, err := client.ConsultaCep(cep)

		//Assert
		assert.NoError(t, err)
		assert.Equal(t, cep, dadosCep.Cep)
		assert.Empty(t, dadosCep.Erro)
	})

	t.Run("Consulta Cep inexistente", func(t *testing.T) {
		//Arrange
		client := NewViaCepClient()
		cep := "00000000"

		//Act
		_, err := client.ConsultaCep(cep)

		//Assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, erros.ErrZipCodeNotFound)
	})

	t.Run("Consulta Cep invalido", func(t *testing.T) {
		//Arrange
		client := NewViaCepClient()
		cep := "058931300"

		//Act
		_, err := client.ConsultaCep(cep)

		//Assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, erros.ErrInvalidZipCode)

		cep = "089313a0"

		//Act
		_, err = client.ConsultaCep(cep)

		//Assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, erros.ErrInvalidZipCode)
	})
}
