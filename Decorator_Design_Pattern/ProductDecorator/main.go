package main

import "fmt"

func main(){
	phone:=&Phone{}


	backcase:=&BackCase{
		caseAccessories: phone,
	}

	tmeper1:=&Temper{
		temper1: backcase,
	}

	fmt.Println(tmeper1.getPrice())
}