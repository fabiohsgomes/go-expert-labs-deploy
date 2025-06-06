package clients

import (
	"testing"

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
		assert.Equal(t, cep, dadosCep.Cep, "CEP should match the requested CEP")
	})

	t.Run("Consulta Cep incorreto", func(t *testing.T) {
		//Arrange
		client := NewViaCepClient()
		cep := "00000000"

		//Act
		dadosCep, err := client.ConsultaCep(cep)

		//Assert
		assert.NoError(t, err)
		assert.Empty(t, dadosCep.Cep)
		assert.NotEmpty(t, dadosCep.Erro)
	})

	t.Run("Consulta Cep invalido", func(t *testing.T) {
		//Arrange
		client := NewViaCepClient()
		cep := "058931300"

		//Act
		_, err := client.ConsultaCep(cep)

		//Assert
		assert.Error(t, err)
	})
}
