package main

import "fmt"



func main() {
	hotel := &hotelTicketBooking{HotelName: "ITC Grand Chola", location: "Banglore", star: "five Star"}
	car := &CarRental{CarName: "BMW 7 Series", carNumber: "KN 0A 2021", variant: "Diesel Variant"}
	flight := &flightBooking{Name: "Air Arabia", Model: "Boeing 777", FromDestination: "Chennai", ToDestination: "Banglore"}


	
	main:= packedTicket(flight, hotel, car)
	fmt.Println(main)

	main1:=CancelTicket()
	fmt.Println(main1)
	

   

	
}
