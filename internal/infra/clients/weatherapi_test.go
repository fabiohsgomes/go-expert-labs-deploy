package clients

import (
	"testing"

	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestConsultaClima(t *testing.T) {
	config.LoadConfig("../../..")

	t.Run("Consulta Clima com sucesso", func(t *testing.T) {
		// Arrange
		client := NewWeatherApiClient()
		cidade := "São Paulo"

		// Act
		weatherResponse, err := client.ConsultaClima(cidade)

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, weatherResponse.Location.Region, "Location name should not be empty")
	})
}
