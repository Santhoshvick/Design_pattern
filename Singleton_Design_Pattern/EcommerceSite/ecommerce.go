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




func (p *ProductCard)AddProduct(userName string,Password string)*ProductCard{
	
	if(instance==nil){
		once.Do(func(){
			user:="santhosh"
			password:="1234"
			if(user==userName&&Password==password){
				instance=&ProductCard{}
				val=true
			   fmt.Println("Login  Successfully")
			}else{
				fmt.Println("Login Failed")
				
			}
		})
	}

	
    if val {
		instance:=&ProductCard{ProductId: p.ProductId,ProductName: p.ProductName,ProductPrice: p.ProductPrice,ProductDescription: p.ProductDescription}
	   return instance
	}

	return nil
	
}