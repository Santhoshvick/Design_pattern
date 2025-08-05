package main

import "fmt"

func main(){

	kia,_:=getCar("Kia")
	fmt.Println(kia)
	hyundai,_:=getCar("Hyundai")
	fmt.Println(hyundai)

	printCar(kia)
	printCar(hyundai)
}

func printCar(c Icar){
	fmt.Printf("Car %s",c.getName())
	fmt.Println()
	fmt.Printf("car Price %d",c.getPrice())
	fmt.Println()
}