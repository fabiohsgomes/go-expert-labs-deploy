package erros

import (
	"errors"
)

var ErrInvalidZipCode = errors.New("invalid zipcode")
var ErrZipCodeNotFound = errors.New("can not find zipcode")
