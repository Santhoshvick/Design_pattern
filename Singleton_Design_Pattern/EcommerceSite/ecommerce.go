package main

import (
	"fmt"
	"sync"
)

type Login struct{
	Username string
	Password string
}


type ProductCard struct{
	Login
	ProductId int64
	ProductName string
	ProductPrice int64
	ProductDescription string
}

var instance *ProductCard

var once sync.Once
var val=false




func AddProduct(userName string,Password string)*ProductCard{
	
	if(instance==nil){
		once.Do(func(){
			user:="santhosh"
			password:="1234"
			if(user==userName&&Password==password){
				instance=&ProductCard{ProductId: 9908,ProductName:"Mackbook Air",ProductPrice: 120000,ProductDescription: "Macbook air with 14 inch display"}
				val=true
			   fmt.Println("Login  Successfully")
			}else{
				fmt.Println("Login Failed")
				
			}
		})
	}
    if val {
	   return instance
	}

	return nil
}

func(p *ProductCard)Update(pid int64,pname string,price int64,pdescription string){
	p.ProductId=pid
	p.ProductName=pname
	p.ProductPrice=price
	p.ProductDescription=pdescription

	fmt.Print("Product Id:",p.ProductId)
	fmt.Print("Product Name:",p.ProductName+"\t")
	fmt.Print("Product Price:",p.ProductPrice)
	fmt.Print("Product Description:",p.ProductDescription)
}