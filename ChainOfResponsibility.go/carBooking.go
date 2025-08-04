package main

import "fmt"

type CarBooking struct{
	next CarDepartment
}

func (b *CarBooking)execute(c *Customer){
	if c.Booked{
          fmt.Println("The Car is Booked Successfully")
		  b.next.execute(c)
		  return
	}
	fmt.Println("Customer is waiting to book the car")
	c.Booked=true
	b.next.execute(c)
}

func (b *CarBooking)setNext(next CarDepartment){
	b.next=next
}