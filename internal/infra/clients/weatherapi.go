package clients

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/config"
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/erros"
)

type WeatherApiClient struct {
	key string
}

var weatherapiuri = "https://api.weatherapi.com/v1/current.json"

func NewWeatherApiClient() *WeatherApiClient {
	cfg := config.Get()

	return &WeatherApiClient{
		key: cfg.GetWeatherApiKey(),
	}
}

func (c *WeatherApiClient) ConsultaClima(cidade string) (*WeatherResponse, error) {
	weatherResponse := &WeatherResponse{}
	weatherErrorResponse := WeatherErrorResponse{}

	req, err := http.NewRequest("GET", weatherapiuri, nil)
	if err != nil {
		return weatherResponse, err
	}

	req.Header.Set("Accept", "application/json")
	url := req.URL
	q := url.Query()
	q.Set("q", cidade)
	q.Set("lang", "pt")
	q.Set("key", c.key)
	url.RawQuery = q.Encode()

	req.URL = url

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return weatherResponse, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		json.Unmarshal(body, &weatherErrorResponse)

		if weatherErrorResponse.ErrorCode() == 1006 {
			return weatherResponse, erros.ErrCityNotFound
		}

		return weatherResponse, weatherErrorResponse
	}

	json.Unmarshal(body, &weatherResponse)

	return weatherResponse, nil
}
