package main

import "fmt"


type ATM interface{
    process(*Click)
}


type Click struct{
	atm ATM
}

func(c *Click)setState(atm1 ATM){
	c.atm=atm1
}

func(c *Click)process(){
	c.atm.process(c)
}

type CardInserted struct{}

func(c *CardInserted)process(c1 *Click){
	fmt.Println("The Card is Inserted Successfully")
	fmt.Println("Please Enter the 4 digit Card pin")
	c1.setState(&PinEntered{})
}

type PinEntered struct{}

func(p *PinEntered)process(c *Click){
	fmt.Println("Please Enter the Amount To be Withdraw")
	c.setState(&Transaction{})
}


type Transaction struct{}


func(t *Transaction)process(c *Click){
	fmt.Println("The Transaction is Completed successfully")
	c.setState(nil)
}

func main(){
	click:=&Click{}

	click.setState(&CardInserted{})

	click.process()
	pin:=0
	fmt.Scanf("%d",&pin)
	if pin==0{
		fmt.Println("Please Enter the Valid 4 digit pin")
	}
	click.process()
	amount:=0
	fmt.Scanf("%d",&amount)
	if(amount<=0){
		fmt.Println("Please Enter the valid amount to be withdrawed")
	}
	fmt.Printf("The amount %d withdraw successfully",amount)
	fmt.Println()
	click.process()
}

