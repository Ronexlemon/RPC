package server

import (
	"log"
	"net"

	grpc "google.golang.org/grpc"
	pb "google.golang.org/protobuf/proto"
	
)

type Server struct{
	pb.Message
}
var addr string = "127.0.0.1:9090"

func MainServer(){
	lis,err := net.Listen("tcp",addr)
	if err !=nil{
		log.Fatal("Failed \n",err)
	}
	log.Printf("Listening %s ........ \n",addr)
	

	s:=grpc.NewServer()
	//pb.MessageName(s Server{})
	if err = s.Serve(lis); err !=nil{
		log.Fatal("Failed \n",err)

	}
}