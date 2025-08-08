package main

import (

	"fmt"
	"sync"
)

type status string
const (
	Booked status="booked"
	blocked status="blocked"
)
	


type SeatManager struct{
	seatNo string
	NoOfSeatsBooked int64
	status status
}

var instance *SeatManager

var lock=sync.Mutex{}

func BookSeat()*SeatManager{
	if(instance==nil){
		lock.Lock()
		defer lock.Unlock()
		if(instance==nil){
			fmt.Println("The Seats of the theatre  is Been managed by Seat Manager ")
			instance=&SeatManager{seatNo: "13A-13D",NoOfSeatsBooked: 4,status: status(Booked)}
		}
	}
	if(instance.status=="blocked"){
		fmt.Println ("Your Ticket Has Been Blocked Due to Payment Failed")
	}

	return instance
}


func (s *SeatManager)Update(sno string,seatBooked int64,status status){
	s.seatNo=sno
	s.NoOfSeatsBooked=seatBooked
	s.status=status

	fmt.Print("Seat No:",s.seatNo)
	fmt.Print("No Of Seats Booked:",s.NoOfSeatsBooked)
	fmt.Print("Status:",s.status)
}