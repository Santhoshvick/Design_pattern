package main

import "fmt"


func main(){
	r1:=&BasicRoom{}

	room1:=r1.addWifi("20MBPS").addBreakfast("IDly with sambar and Chutney").addSeaView("").addLateCheckOut("").build()
	fmt.Println(room1)
}