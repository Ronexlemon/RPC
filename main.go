package main

import (
	itemserver "gorpc/basic/server"
	itemserver2 "gorpc/basic/client"
	// "gorpc/client"
	// "gorpc/server"
)

func main(){
	// server.Server()
	// client.Client()
	itemserver.ItemServer()
	itemserver2.ClientBasic()
	
}