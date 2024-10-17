package main

import ( arithmeticserver "gorpc/arithmetic/server"
arithmeticclient "gorpc/arithmetic/client"

)

// itemserver "gorpc/basic/server"
// itemserver2 "gorpc/basic/client"
// "gorpc/client"
// "gorpc/server"

func main(){
	// server.Server()
	// client.Client()
	// itemserver.ItemServer()
	// itemserver2.ClientBasic()
	go arithmeticserver.ArithServer()
	arithmeticclient.ArithClient()
	
}