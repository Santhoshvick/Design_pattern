package main

import "fmt"

func main(){
	cash:=&Cash{}
	cash.makePayment(20000)

	fmt.Println("-------------------------------------------")

	cc:=&CreditCard{}
	cc.makePayment(4000.89)
	fmt.Println()

	dc:=&DebitCard{}
	dc.makePayment(10000.23)
}