package main

import "fmt"

type HotelBooking interface{
	book(*Room)
}

type Room struct{
	hotelBooking HotelBooking
}

func(r *Room)setState(room HotelBooking){
	r.hotelBooking=room
}

func(r *Room)book(){
	r.hotelBooking.book(r)
}

type Available struct{
	Availability bool
}

func(a *Available)book(r *Room){
	a.Availability=true
	if(a.Availability){
		fmt.Println("The Room is Available you can start your booking process")
		r.setState(&Booked{})
	}
}

type Booked struct{
	Name string
	IdProof string
	TotalPersons int32
	NoOfDays int32
}

func(b *Booked)book(r *Room){
	fmt.Printf("The Room for the customer %s has been Booked successfully for %d days",b.Name,b.NoOfDays)
	r.setState(&CheckIn{})
}

type CheckIn struct{}

func(c *CheckIn)book(r *Room){
	fmt.Println("Welcome to XYZ Resort this is Your key enjoy your day")
	r.setState(&CheckOut{})
}

type CheckOut struct{}

func(c *CheckOut)book(r *Room){
	fmt.Println("Thanks for Choosing our resort ,Please Visit Again!!!...")
	r.setState(nil)
}


func main(){
	room:=&Room{}
	room.setState(&Available{})

	cust1:=&Booked{Name:"john",IdProof: "Aadhar Card",NoOfDays: 3,TotalPersons: 4}

	room.book()
	cust1.book(room)
	room.book()
	room.book()
}