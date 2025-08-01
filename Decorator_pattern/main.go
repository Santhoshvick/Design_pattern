package main

import "fmt"

func main(){
	veg:=&vegMania{}
	

	chessyToppings:=&CheesyPizza{
		pizza: veg,
	}

	TomatoTopping1:=&TomatoTopping{
		pizza: chessyToppings,
	}
	fmt.Printf("The total cost of the pizza with extra topping is %v",TomatoTopping1.getPrice())
}

