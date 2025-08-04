package main

import "fmt"

type PackageTicketBooking struct {
	fbooking *flightBooking
	hbooking *hotelTicketBooking
	cbooking *CarRental
}

func packedTicket(fbook *flightBooking, hbook *hotelTicketBooking, cbook *CarRental) *PackageTicketBooking {
	fmt.Println(fbook.Bookticket())
	fmt.Print(hbook.BookHotel())
	fmt.Println(cbook.BookCar())
	return &PackageTicketBooking{
		fbooking: fbook,
		hbooking: hbook,
		cbooking: cbook,
	}
}

func CancelTicket() *PackageTicketBooking {
	fmt.Println("Ticket is canceled Successfully")
	return &PackageTicketBooking{
		fbooking: nil,
		hbooking: nil,
		cbooking: nil,
	}
}



func (s *PackageTicketBooking) bookPackage() {
	s.cbooking.BookCar()
	s.hbooking.BookHotel()
	s.hbooking.BookHotel()
}
