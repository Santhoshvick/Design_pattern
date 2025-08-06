package main

import "fmt"

type CreditCard struct{
	Amount float64
}

func (cc *CreditCard)makePayment(amount float64){
	cc.Amount=amount
	fmt.Printf("The Payment method is creditcard and the bill amount is %v",amount)
}