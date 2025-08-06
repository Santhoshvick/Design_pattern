package main

import "fmt"

type BankAccountFacade struct{
	account *Account
	ledger *Ledger

}


func (b *BankAccountFacade)withdraw(amount float64){
	if b.account.getBalance(amount){
		b.account.debit(amount)
		fmt.Println(b.account.debit(amount))  
		b.ledger.Record("withdraw successfull",amount)
	}
}