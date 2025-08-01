package main

import "fmt"



func main(){
	flight1,_:=getflight("Emirates")
	fmt.Println(flight1)
	printFlight(flight1)

}




func printFlight(f Iflight){
	fmt.Printf("Flight Name :%s",f.getFlightName())
	fmt.Println()
	fmt.Printf("Flight Price: %d",f.getFlightTicketPrice())
	fmt.Println()
	fmt.Printf("Available Class :%s",f.getFlightClass())

}
