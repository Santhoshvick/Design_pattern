package main

import "fmt"

func getflight(name string)(Iflight,error){
	if(name=="Emirates"){
		return  Emirates1(),nil
	}
	return nil,fmt.Errorf("Flight does not Found %v")
}