package main

import "fmt"

type DebitCard struct{
	Payment
	Amount float64
}

func(c *DebitCard)makePayment(amount float64){
	c.Amount=amount
	fmt.Printf("The Payment method is Debit Card and the bill amount is %v",amount)
	fmt.Println()
	fmt.Println("Debit Card Payment Done  successfully")
}