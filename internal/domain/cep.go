package domain

import (
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/erros"
	"github.com/fabiohsgomes/go-expert-labs-deploy/internal/helpers"
)

type Cep struct {
	codigo string
}

func NewCep(codigo string) (*Cep, error) {
	cep := &Cep{
		codigo: helpers.NormalizeZipCode(codigo),
	}

	if !cep.validar() {
		return nil, erros.ErrInvalidZipCode
	}

	return cep, nil
}

func (c *Cep) Codigo() string {
	return c.codigo
}

func (c Cep) validar() bool {
	return len(c.codigo) == 8
}
