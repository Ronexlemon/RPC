package main

import (
	"gorpc/client"
	"gorpc/server"
)

func main(){
	server.Server()
	client.Client()
}