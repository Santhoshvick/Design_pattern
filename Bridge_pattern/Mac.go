package main

import (
	"fmt"
)
type mac struct{
	printer Printer
}

func (m *mac)Print(){
	fmt.Println("Printer request for mac")
	m.printer.Printfile()
}

func (m *mac)setPrinter(p Printer){
	m.printer=p
}

