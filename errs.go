package gosrisdk

import (
	"errors"
)

var (
	ErrRucVacio        = errors.New("el RUC esta vacio")
	ErrRucNoValido     = errors.New("el RUC no es valido")
	ErrResponseNil     = errors.New("el response del SRI es nil")
	ErrRucNoRegistrado = errors.New("el RUC no esta registrado")
)
