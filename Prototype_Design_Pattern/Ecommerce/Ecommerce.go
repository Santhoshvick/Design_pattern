package main

import "fmt"

type Cloneable interface{
	clone()Cloneable
}
type Product struct{
	ProductId int64
	ProductName string
	ProductPrice string
	ProductVariant string
}

func(p *Product)clone()Cloneable{
	return &Product{
		ProductId: p.ProductId,
		ProductName: p.ProductName,
		ProductPrice: p.ProductPrice,
		ProductVariant: p.ProductVariant,
	}
}

func main(){
	product1:=&Product{ProductId: 1,ProductName: "Samsung Galaxy s24 ultra",ProductPrice:"120000",ProductVariant: "8GB RAM 128GB Storage"}
	fmt.Println("Product1:",product1)
	product2:=product1.clone().(*Product)
	product2.ProductVariant="12GB RAM 256GB Storage"
	fmt.Println("Product2:",product2)
}
