package itemserver

import (
	"fmt"
	itemserver "gorpc/basic/server"
	"log"
	"net/rpc"
)


func ClientBasic(){
	var reply itemserver.Item
	var db []itemserver.Item
	
	 	item := itemserver.Item{Name: "Golang",Price: 100.00}
	item2 := itemserver.Item{Name: "Python",Price: 200.00}
	item3 := itemserver.Item{Name: "Java",Price: 300.00}
	itemForEdit := itemserver.Item{Name: "Solidity",Price: 500.00}
	//itemForEdit2 := itemserver.Item{Name: "Solidity",Price: 500.00}

	client,err := rpc.DialHTTP( "tcp", "localhost:9090")
	fmt.Println("dialing .......")
	fmt.Println("dialing .......")
	if err != nil {
		log.Fatal("Error dialing",err)}
	defer client.Close()
	 client.Call("API.CreateItem",&item,&reply)
	 client.Call("API.CreateItem",&item3,&reply)
	 client.Call("API.CreateItem",&item2,&reply)
	 client.Call("API.CreateItem",&item2,&reply)
	 client.Call("API.GetDB","",&db)
	 fmt.Println("database",db)
	 client.Call("API.EditItem",&itemForEdit,&reply)
	 client.Call("API.GetDB","",&db)
	 fmt.Println("database after Edit",db)

	 client.Call("API.GetItemByName",item2.Name,&reply)
	 fmt.Println("Item",reply)

	 

}