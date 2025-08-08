package main

import (
	"fmt"
	"sync"
)

type Displayer struct{
	ItemName string
	ItemPrice float64
	ItemDescription string
}

var instance *Displayer

var lock=sync.Mutex{}

func show()*Displayer{
	if(instance==nil){
		lock.Lock()
		defer lock.Unlock()
		if (instance==nil){
			instance=&Displayer{ItemName: "Veg Meals",ItemPrice: 250.40,ItemDescription: "Veg Comboi Meal"}
			fmt.Println("Display is ready To display the Food Items")
		}
	}

	return instance
}

func(d *Displayer)update(description string,name string,price float64){
	d.ItemDescription=description
	d.ItemName=name
	d.ItemPrice=price

	fmt.Print("Item Name:",d.ItemName)
	fmt.Print("Item Price",d.ItemPrice)
	fmt.Print("Item Description",d.ItemDescription)
}

