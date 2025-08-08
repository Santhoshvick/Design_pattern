package main

import "fmt"


type Order struct{
	next Ecommerce
}

func(o *Order)execute(c *Customer){
	fmt.Println("The Product is Yet to be Ordered please complete the Payment first")
	o.next.execute(c)
}


func(o *Order)setNext(next Ecommerce){
	o.next=next
}