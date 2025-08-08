package main

import "fmt"

func main(){
	Np:=&NormalPc{}

	pc1:=Np.addMonitor("Samsung").addKeyboard("Logitech").addMouse("Logitech").addCpu("512GB","16GB RAM","RTX 50 series (future):")
	fmt.Println(pc1)
}