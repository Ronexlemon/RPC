package arithmeticserver

import (
	"context"

	arithmeticproto "gorpc/arithmetic/proto"
)

type Server struct{}

func (s *Server) Add(ctx context.Context, request *arithmeticproto.Request) (*arithmeticproto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &arithmeticproto.Response{Result: result}, nil

}

func (s *Server) Multiply(ctx context.Context, request *arithmeticproto.Request) (*arithmeticproto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &arithmeticproto.Response{Result: result}, nil
}

func ArithServer() {

}
