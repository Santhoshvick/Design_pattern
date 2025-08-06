package main

import "fmt"

func main(){
	veg:=&vegMania{}

	fmt.Println(veg.getPrice())


	chessyToppings:=&CheesyPizza{
		pizza: veg,
	}
	fmt.Println(chessyToppings.getPrice())

	TomatoTopping1:=&TomatoTopping{
		pizza: chessyToppings,
	}
	fmt.Printf("The total cost of the pizza with extra topping is %v",TomatoTopping1.getPrice())
}

