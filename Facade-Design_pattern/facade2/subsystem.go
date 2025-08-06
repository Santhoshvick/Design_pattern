package main

import "fmt"

type flightBooking struct{
	Name string
	Model string
	FromDestination string
	ToDestination string
}


func (f *flightBooking)Bookticket() *flightBooking{
	fmt.Println("Flight Ticket is Booked successfully")
	fmt.Println("------------------------------")
	return &flightBooking{
		Name: f.Name,
		Model: f.Model,
		FromDestination: f.FromDestination,
		ToDestination: f.ToDestination,
	}
}


type hotelTicketBooking struct{
	HotelName string
	location string
	star string
}

func (h *hotelTicketBooking)BookHotel()*hotelTicketBooking{
	fmt.Println("The Hotel Ticket is booked successfully")
	fmt.Println("------------------------------")
	return &hotelTicketBooking{
		HotelName: h.HotelName,
		location: h.location,
		star: h.star,
	}
}

type CarRental struct{
	CarName string
	carNumber string
	variant string
}

func (c *CarRental)BookCar()*CarRental{
	fmt.Println("------------------------------------")
	fmt.Println("The car is booked successfully")
	
	return &CarRental{
		carNumber: c.carNumber,
		CarName: c.CarName,
		variant: c.variant,
	}
}