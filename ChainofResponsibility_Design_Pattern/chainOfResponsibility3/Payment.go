package main

import "fmt"

type Payment struct{
	next Ecommerce
}


func(p *Payment)execute(c *Customer){
	fmt.Println("The Payment is Completed Successfully")
	p.next.execute(c)
}

func(p *Payment)setNext(next Ecommerce){
	p.next=next
}