package main

import "fmt"

type Delivery struct{
	next Ecommerce
}

func(d *Delivery)execute(c *Customer){
	fmt.Println("The Product is Delivered to your Address successfully",c.CustomerAddress)
	d.next.execute(nil)
}

func(d *Delivery)setNext(next Ecommerce){
	d.next=nil;
}