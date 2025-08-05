package main

import "fmt"

type Cloning interface{
	clone()Cloning
}


type Product struct{
	ProductName string
	ProductPrice float64
	ProductId int64
	ProductDescription string
}

func(p *Product)clone()Cloning{

	return &Product{
		ProductName: p.ProductName,
		ProductPrice: p.ProductPrice,
		ProductId: p.ProductId,
		ProductDescription: p.ProductDescription,
	}	
}

func main(){
	product1:=&Product{ProductName: "Macbook pro",ProductPrice: 145000.99,ProductId: 101,ProductDescription: "Latest Generation Macbook pro m4 chip"}
	fmt.Println(product1)

	product2:=product1.clone().(*Product)
	product2.ProductDescription="Old Genereation M1 Chip"
	fmt.Println(product2)
}