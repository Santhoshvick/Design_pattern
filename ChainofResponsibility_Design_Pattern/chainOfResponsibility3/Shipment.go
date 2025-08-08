package main

import "fmt"


type Shipment struct{
	next Ecommerce
}

func(s *Shipment)execute(c *Customer){
	fmt.Println("The Product has been dispatched from WareHouse Successfully")
	s.next.execute(c)
}

func(s *Shipment)setNext(next Ecommerce){
	s.next=next
}