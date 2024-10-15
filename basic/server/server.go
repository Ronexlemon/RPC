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

func GetItemByName(name string)(*Item,error){
	
	for _,value := range Database{
		if value.Name == name{
			return &value, nil
	}}
	return nil, fmt.Errorf("item with name %s not found", name)
}
func CreateItem(item *Item) (*Item, error) {
	Database = append(Database, *item)  
	return item, nil
}

func EditItem(name string,item *Item)(*Item,error){
	for index,value := range Database{
		if value.Name == name{
			
			Database[index] = *item
			return item, nil
	}}
	return nil, fmt.Errorf("item with name %s not found", name)

}

func ItemServer(){
	fmt.Println("Server is running for Items")
	item := Item{Name: "Golang",Price: 100.00}
	item2 := Item{Name: "Python",Price: 200.00}
	item3 := Item{Name: "Java",Price: 300.00}
	itemForEdit := Item{Name: "Solidity",Price: 500.00}
	CreateItem(&item)
	CreateItem(&item2)
	CreateItem(&item3)
	fmt.Println("Database Items",Database)

	editItem,err := EditItem("Java",&itemForEdit)
	if err != nil{
		fmt.Println(err)}
		if editItem != nil{
			fmt.Println("The Edited Item is",editItem)}
		fmt.Println("Database Items after edit",Database)
itemnew,err := GetItemByName("Golang")
if err != nil{
	fmt.Println(err)}
	if itemnew != nil{
		fmt.Println("The Item is",itemnew)}

		fmt.Println("Database Items",Database)
		

}