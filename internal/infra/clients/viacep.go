package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/erros"
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/helpers"
)

type ViaCepClient struct {
}

var viacepuri = "https://viacep.com.br/ws/"
var format = "/json/"

func NewViaCepClient() *ViaCepClient {
	return &ViaCepClient{}
}

func (c *ViaCepClient) ConsultaCep(cep string) (*DadosCepResponse, error) {
	if !helpers.ValidateZipCode(cep) {
		return nil, erros.ErrInvalidZipCode
	}

	dadosCep := &DadosCepResponse{}

	req, err := http.NewRequest("GET", viacepuri+cep+format, nil)
	if err != nil {
		return dadosCep, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return dadosCep, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return dadosCep, fmt.Errorf("error fetching data: %s", resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)

	json.Unmarshal(body, &dadosCep)

	if len(dadosCep.Erro) > 0 {
		return dadosCep, erros.ErrZipCodeNotFound
	}

	dadosCep.Cep = helpers.NormalizeZipCode(dadosCep.Cep)

	return dadosCep, nil
}
