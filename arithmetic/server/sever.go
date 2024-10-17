package arithmeticserver

import (
	"context"
	"net"

	arithmeticproto "gorpc/arithmetic/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	listner,err := net.Listen("tcp", ":9090")
	if err !=nil{
		panic(err)
	}
	srv := grpc.NewServer()
	arithmeticproto.RegisterAddServiceServer(srv,&Server{})
	reflection.Register(srv)

	if err = srv.Serve(listner); err != nil{
		panic(err)
	}
	

}
