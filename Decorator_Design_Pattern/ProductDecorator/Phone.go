package main

import "fmt"

type Phone struct{
	
}

func(p *Phone)getPrice()float64{
	fmt.Println("The Phone is Ordered successfully you need any extra accessories for the phone you ordered")
	return 25000.90
}