package server

import "errors"

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (a *Arith) Multiply(ar *Args, result *int) error {
	*result = ar.A * ar.B
	return nil
}

func (a *Arith) Divide(ar *Args, quo *Quotient) error {
	if ar.B == 0 {
		return errors.New("Divide By Zero")
	}
	quo.Quo = ar.A / ar.B
	quo.Rem = ar.A % ar.B

	return nil

}

func (a *Arith) Add(ar *Args, result *int) error {

	*result = ar.A + ar.B
	return nil
}
