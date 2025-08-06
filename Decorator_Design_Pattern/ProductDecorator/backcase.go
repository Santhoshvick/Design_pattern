package main

import "fmt"

type BackCase struct{
	caseAccessories OrderedProduct
}


func(b *BackCase)getPrice()float64{
	fmt.Println("Backcase is added extra with your phone")
	price:=b.caseAccessories.getPrice()
	return price+200.60
}