package main

import "fmt"

type TomatoTopping struct{
	pizza Ipizza
}


func (t *TomatoTopping)getPrice()int64{
	fmt.Println("Extra topping has beeen added to the veg mania pizza")
	price:=t.pizza.getPrice()
	return int64(price)+20
}