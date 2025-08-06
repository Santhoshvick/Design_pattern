package main

import "fmt"

func main(){
	atm:=&ATM{Balance: 20000}
	atm.checkbalance()

	withdraw:=atm.withdraw(50000)
	fmt.Println()
	fmt.Println(withdraw)
}