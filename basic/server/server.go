package server

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

}