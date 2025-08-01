package main

import "fmt"

type CheesyPizza struct{
	pizza Ipizza
}

func(c *CheesyPizza)getPrice()int64{
	fmt.Println("Chessy Topping has been added extra to veg mania")
	price:=c.pizza.getPrice()
	return  int64(price)+50
}

