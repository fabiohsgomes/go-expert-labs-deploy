package clients

import (
	"encoding/json"
	"io"
	"net/http"
)

type WeatherApiClient struct {
}

var weatherapiuri = "https://api.weatherapi.com/v1/current.json"
var key = "da26cd9b6c624664977234238250506"

func NewWeatherApiClient() *WeatherApiClient {
	return &WeatherApiClient{}
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
	q.Set("lang", "pt-br")
	q.Set("key", key)
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
		return weatherResponse, weatherErrorResponse
	}

	json.Unmarshal(body, &weatherResponse)

	return weatherResponse, nil
}
