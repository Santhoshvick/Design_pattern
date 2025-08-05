package main

type onCommand struct{
    device Device
}

func (o *onCommand) execute(){
	o.device.On()
}