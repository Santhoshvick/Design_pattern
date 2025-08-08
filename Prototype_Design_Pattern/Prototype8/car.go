package main

import "fmt"

type clonable interface{
	clone()clonable
}

type Car1 struct{
	CarBrand string
	CarModel string
	CarVariant string
	CarPrice float64
}

func(c *Car1)clone()clonable{
	return &Car1{
		CarBrand:c.CarBrand ,
		CarModel:c.CarBrand ,
		CarVariant: c.CarVariant,
		CarPrice: c.CarPrice,
	}
}

func main(){
	car1:=&Car1{CarBrand:"Range Rover",CarModel:"Evoque",CarVariant:"HSX-Base",CarPrice: 25000000}

	car2:=car1.clone().(*Car1)
	car2.CarVariant="HSY-Mid"
	car2.CarPrice=30000000

	fmt.Println("Car1",car1)
	fmt.Println("Car2",car2)

}