package main

import "fmt"

type Train struct{
	TrainName string
	TrainNumber int64
	FromLocation string
}


func(t *Train)ChooseDestination(destination string){
	fmt.Printf("The Selected Destination is %s",destination)
}

func(t *Train)TravelMode(){
	fmt.Println("The Travel Mode you have selected is First class")
}

func(t *Train)Book(){
	fmt.Println("Confirm Booking")
}

func(t *Train)Confirm(){
	fmt.Printf("TrainName:%s\n",t.TrainName)
	fmt.Printf("TrainNumber: %v\n",t.TrainNumber)
	fmt.Printf("FromLocation :%s\n",t.FromLocation)
}