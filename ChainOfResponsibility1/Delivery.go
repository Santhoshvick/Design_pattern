package main

import "fmt"


type Delivery struct{
	next CarDepartment
}


func (d *Delivery)execute(c *Customer){
	if c.Delivered{
		fmt.Println("the Car Delivered to the customer Successfully!!!!,Happy Motoring")
		d.next.execute(c)
		return
	}
	fmt.Println("The Payment is getting delay so that delivery is also getting delay")
}

func (d *Delivery)setNext(next CarDepartment){
	d.next=next
}