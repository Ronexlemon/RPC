package itemserver

import "fmt"

type Item struct{
	Name string
	Price float64
}

var Database []Item

func GetItems()[]Item{
	return Database
}
type API int

func (a *API) GetItemByName(name string,reply *Item)(error){
	
	for _,value := range Database{
		if value.Name == name{
			*reply = value
			return nil
	}}
	return  fmt.Errorf("item with name %s not found", name)
}
func (a *API) CreateItem(item *Item, reply *Item) ( error) {
	Database = append(Database, *item)  
	*reply = *item
	return nil
}

func (a *API) EditItem(name string,item *Item,reply *Item)(error){
	for index,value := range Database{
		if value.Name == name{
			
			Database[index] = *item
			*reply = *item
			return  nil
	}}
	return  fmt.Errorf("item with name %s not found", name)

}

func ItemServer(){
	fmt.Println("Server is running for Items")
	api :=new(API)
	var item_a Item
	item := Item{Name: "Golang",Price: 100.00}
	item2 := Item{Name: "Python",Price: 200.00}
	item3 := Item{Name: "Java",Price: 300.00}
	itemForEdit := Item{Name: "Solidity",Price: 500.00}
	api.CreateItem(&item,&item_a)
	api.CreateItem(&item2,&item_a)
	api.CreateItem(&item3,&item_a)
	fmt.Println("Database Items",Database)

	err := api.EditItem("Java",&itemForEdit,&item_a)
	if err != nil{
		fmt.Println(err)}
		
		fmt.Println("Database Items after edit",Database)
err = api.GetItemByName("Golang",&item_a)
if err != nil{
	fmt.Println(err)}
	

		fmt.Println("Database Items",Database)
		

}