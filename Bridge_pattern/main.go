package main

import "fmt"

func main(){
	hp:=&Hp{}

	epison:=&Epsion{}

	mac1:=&mac{}
	mac1.setPrinter(epison)
	windows:=&Window{}
	windows.setPrinter(hp)
	mac1.Print()
	fmt.Println()
	windows.Print()


}