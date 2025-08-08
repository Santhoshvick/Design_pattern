package main

import "fmt"

type Payments struct{
	next CarDepartment
}


func (p *Payments)execute(c *Customer){
	if c.PaymentDone{
		fmt.Println("The Payment is done successfully,The car will be delivered to the customer with in 2 days")
		p.next.execute(c)
		return
	}
	fmt.Println("The customer is Waiting for EMI Approval")
	c.PaymentDone=true
	p.next.execute(c)
}

func (p *Payments)setNext(next CarDepartment){
	p.next=next
}