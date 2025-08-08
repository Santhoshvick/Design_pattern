package main

import (
	"fmt"
	"time"
)

type Ecommerece interface{
	order(*Action)
}

type Action struct{
	ecommerce Ecommerece
	OrderId int64
	ProductName string
}


func(a *Action)setState(ecomm Ecommerece){
	a.ecommerce=ecomm
}

func(a *Action)order(){
	a.ecommerce.order(a)
}

type Paid struct{}

func(p *Paid)order(a *Action){
	fmt.Println("The Amount for the product is paid successfully")
	a.setState(&Shipped{})
}

type Shipped struct{}


func(s *Shipped)order(a *Action){
	
	cancel:=true
	if(cancel){
		time.Sleep(2*time.Second)
		a.setState(&Cancel{})
	}else{
	time.Sleep(2*time.Second)
	a.setState(&Delivered{})
	fmt.Println("The Product has been shipped to the particular delivery Address")
	}
}

type Review struct{}
func(r *Review)order(a *Action){
	time.Sleep(2*time.Second)
	fmt.Println("Please Share your feedback of the product you received through the google form url:https://docs.google.com/forms/d/1DKSvXdZNzViDUfixlKHj28YUYm1fBGe8KxspusYY6K8/edit")
	a.setState(nil)
}

type Cancel struct{}

func(c *Cancel)order(a *Action){
	time.Sleep(2*time.Second)
	fmt.Println("The Cancellation request has been raised successfully")
	a.setState(&Refund{})
}

type Delivered struct{}

func(d *Delivered)order(a *Action){
	time.Sleep(2*time.Second)
	action:=&Action{ProductName: a.ProductName,OrderId: a.OrderId}
	fmt.Println("The Order Details:",action)
	fmt.Println("The Item is Delivered Successfully")
	a.setState(&Review{})
}



type Refund struct{}

func(r *Refund)order(a *Action){
	time.Sleep(2*time.Second)
	fmt.Printf("The Cancellation for the product %v has been approved Successfully %v",a.ProductName,a.OrderId)
	a.setState(nil)
}

func main(){
	a:=&Action{ProductName: "Samsung Galaxy S24 Ultra",OrderId: 102012}
	
	a.setState(&Paid{})

	a.order()
	a.order()
	a.order()
	a.order()


	
}

