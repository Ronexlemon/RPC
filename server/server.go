package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

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

func Server() {
	arith := new(Arith)

	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	go http.Serve(l, nil)

}
