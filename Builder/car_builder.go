package main

import "fmt"

type IcarBuilder interface{
	setColor(color1 string)
	setTyreSize(tyre1 string)
	setEngine(engineName string)
	setModel(model string)
	getCar() Car
}


func buildcar(brand string)(IcarBuilder,error){
	if brand=="Kia"{
		return newKia(),nil
	}else if brand=="Hyundai"{
		return newHyundai(),nil
	}
	return nil,fmt.Errorf("Wrond Model Type")
}




