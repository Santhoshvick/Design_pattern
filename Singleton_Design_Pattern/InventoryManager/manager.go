package main

import (
	"fmt"
	"sync"

	"google.golang.org/genproto/googleapis/type/date"
)

type InvertoryManager struct{
	ProductName string
	ProductQuantity string
	ProductExpiryDate date.Date
}

var details *InvertoryManager

var lock=sync.Mutex{}

func getInstance()*InvertoryManager{
	if(details==nil){
		lock.Lock()
		defer lock.Unlock()
		if (details==nil){
			fmt.Println("The Inventory is created successfully")
	     	details=&InvertoryManager{ProductName: "demo",ProductQuantity:"demo1",ProductExpiryDate: date.Date{Year: 0}}
		}
	}
	return details
}


func(i *InvertoryManager)update(productname string,productQuantity string,productExpiry date.Date){
	i.ProductName=productname
	i.ProductQuantity=productQuantity
	i.ProductExpiryDate=productExpiry
	fmt.Print("Product Name:",i.ProductName)
	fmt.Print("Product Quantity:",i.ProductQuantity)
	fmt.Print("Product Expiry:",i.ProductExpiryDate.Year)
}