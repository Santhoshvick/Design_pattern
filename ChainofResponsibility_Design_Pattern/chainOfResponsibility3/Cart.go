package main

import "fmt"

type Cart struct{
	next Ecommerce
}

type Product struct{
	ProductName string
	ProductPrice float64
	ProductQuantity int64
}

func(c1 *Cart)execute(c *Customer){
	fmt.Println("The Product is added to the cart")
	c1.next.execute(c)
}

func (c1 *Cart)setNext(next Ecommerce){
	c1.next=next
}