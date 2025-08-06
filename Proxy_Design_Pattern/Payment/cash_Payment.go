package main

import "fmt"

type Cash struct{
	Payment
	Amount float64
}

func(c *Cash)makePayment(amount float64){
	c.Amount=amount
	fmt.Printf("The Payment method is cash and the bill amount is %v",amount)
	fmt.Println()
	fmt.Println("Cash Payment Done  successfully")
}