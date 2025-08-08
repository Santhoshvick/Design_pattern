package main

import "fmt"

func main(){
	
	val1:=BookSeat()
	val1.Update("10E-10F",2,status(Booked))
	fmt.Println()


	val2:=BookSeat()
	val2.Update("1A-1G",7,status(Booked))
	fmt.Println()

}