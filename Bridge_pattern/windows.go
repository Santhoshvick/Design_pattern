package main

import (
	"fmt"
)

type Window struct{
	printer Printer
}


func (w *Window)Print(){
	fmt.Println("Printer request from windows")
	w.printer.Printfile()
}

func (w * Window)setPrinter(p Printer){
	w.printer=p
}
