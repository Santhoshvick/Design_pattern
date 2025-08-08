package main

import "fmt"

type Flight struct{
	FlightName string
	FlightNumber int64
	FromLocation string
}


func(f *Flight)ChooseDestination(destination string){
	fmt.Printf("The Selected Destination is %s",destination)
}

func(f *Flight)TravelMode(){
	fmt.Println("The Travel Mode you have selected is First class")
}

func(f *Flight)Book(){
	fmt.Println("Confirm Booking")
}

func(f *Flight)Confirm(){
	fmt.Printf("FlightName:%s\n",f.FlightName)
	fmt.Printf("FlightNumber: %v\n",f.FlightNumber)
	fmt.Printf("FromLocation :%s\n",f.FromLocation)
}