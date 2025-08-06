package main

import "fmt"

type Account struct{
	AccountNumber int64
	AccountHolderName string
	AvailableBalance float64
}


func (a *Account)getBalance(amount float64) bool{
	val:=false
	account1:=&Account{AccountNumber: 1234567,AccountHolderName: "Santhosh",AvailableBalance: 25000.00}
	if amount>account1.AvailableBalance{
		fmt.Println("Entered amount is greater then your balance,Please check the balance once")
		return val
	}
	fmt.Println("The Withdraw is successfull",amount)
	newBalance:=account1.AvailableBalance-amount
	fmt.Println("The New Balance is:",newBalance)
	return val
}

func(a *Account) debit(amount float64)float64{
	balance:=a.getBalance(amount)
	if balance{
		return amount
	}
	return 0.0
}



type Ledger struct{  
	Message string
	Amount float64

}

func (l *Ledger)Record(message1 string,amount float64)*Ledger{
	return &Ledger{
		Message: message1,
		Amount: amount,
	}


	

}