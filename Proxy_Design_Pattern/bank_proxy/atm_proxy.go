package main

import "fmt"

type ATM struct{
	Balance float32
	Bank
}

func(a *ATM)checkbalance(){
	fmt.Printf("The Available Balance is %v",float32(a.Balance))
}

func(w *ATM)withdraw(amount float64)string{
	if (amount>float64(w.Balance)){
		return "The Withdraw failed Balance less then your withdraw amount"
	}
    
	w.Balance=w.Balance-float32(amount)
	return "Withdraw Successfull"
}

