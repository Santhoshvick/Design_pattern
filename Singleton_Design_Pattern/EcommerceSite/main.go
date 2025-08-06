package main

import "fmt"

func main(){
	ProductCart1:=&ProductCard{ProductId: 10001,ProductName: "Lenovo NoteBook Pro",ProductPrice: 86799,ProductDescription: "Lenovo 14 inch lcd display with storage of 512GB"}
	fmt.Println(ProductCart1.AddProduct("santhosh","1234"))

	productCart2:=&ProductCard{ProductId: 10002,ProductName: "Samsung galaxy s25 Ultra",ProductPrice: 99999,ProductDescription: "Samsung s25 With a Curved display"}
	fmt.Println(productCart2.AddProduct("santhosh","1234"))
}