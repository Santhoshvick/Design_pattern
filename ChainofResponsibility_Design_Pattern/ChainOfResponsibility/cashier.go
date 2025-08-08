package main

import "fmt"

type Cashier struct{
	next Department
}


func (c * Cashier)execute(p *Patient){
	p.CashPayed=true
	if p.CashPayed{
        fmt.Println("The Amount is successfully Received from the patient for the medical consulation")
		c.next.execute(p)
		return
	}
	fmt.Println("Patienet is waiting to pay his/her bill")
}

func (c *Cashier)setNext(next Department){
	c.next=next
}