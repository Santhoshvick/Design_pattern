package main

type offCommand struct{
	device Device
}


func (of *offCommand)execute(){
	of.device.off()
}
