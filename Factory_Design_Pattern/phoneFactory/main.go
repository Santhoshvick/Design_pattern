package main

import "fmt"

func main(){

	phone1:=getPhone("Apple")
	fmt.Println(phone1)

	phone2:=getPhone("Samsung")
	fmt.Println(phone2)

	displayDetails(phone1)
	displayDetails(phone2)
}

func displayDetails(g FPhone){
	fmt.Println("Phone Details")
	fmt.Println("----------------------------------------")
	fmt.Println("Phone Name:",g.getName())
	fmt.Println("Phone Price:",g.getPrice())
	fmt.Println("Phone Model:",g.getModel())
	fmt.Println("Storage:",g.getStorage())
	fmt.Println("============================================")

}